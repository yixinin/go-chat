package logic

import (
	"context"
	"encoding/json"
	"go-chat/cache"
	"go-chat/protocol"
	"go-lib/log"
	"go-lib/registry"
	"go-lib/utils"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type NotifyFunc func(uid string, msg interface{}) (bool, error)

//发送消息/邀请等。。。
type ChatLogic struct {
	sync.RWMutex
	watcher     registry.Watcher
	roomClients map[string]protocol.RoomServiceClient
	Notify      NotifyFunc
	stop        chan bool
}

func NewChatLogic(watcher registry.Watcher, notifyFuncs ...NotifyFunc) *ChatLogic {
	var notifyFunc NotifyFunc = nil
	if len(notifyFuncs) > 0 {
		notifyFunc = notifyFuncs[0]
	}
	var c = &ChatLogic{
		roomClients: make(map[string]protocol.RoomServiceClient, 2),
		Notify:      notifyFunc,
		watcher:     watcher,
	}

	go c.Watch()
	return c
}

func (s *ChatLogic) AddRoomConn(addr string) {
	s.Lock()
	defer s.Unlock()
	var conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	if _, ok := s.roomClients[addr]; ok {
		delete(s.roomClients, addr)
	}
	var client = protocol.NewRoomServiceClient(conn)
	s.roomClients[addr] = client
}
func (s *ChatLogic) DelRoomConn(addr string) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.roomClients[addr]; ok {
		delete(s.roomClients, addr)
	}
}

func (s *ChatLogic) UpdateRoomConn(addr string) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.roomClients[addr]; ok {
		return
	}
	var conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return
	}
	var client = protocol.NewRoomServiceClient(conn)
	s.roomClients[addr] = client
}

func (s *ChatLogic) SendMessage(req *protocol.SendMessageReq, ack *protocol.SendMessageAck) (err error) {

	ack.Header.Code = 200
	ack.Header.Msg = "success"
	return
}

func (s *ChatLogic) RealTime(req *protocol.RealTimeReq, ack *protocol.RealTimeAck) (err error) {
	switch req.Protocol {
	case "tcp":
	case "ws":
	}

	var users []*protocol.RoomUser

	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if req.Uid != "" {
		users = []*protocol.RoomUser{
			&protocol.RoomUser{
				Uid:   req.Header.Uid,
				Token: req.Header.Token,
			},
			&protocol.RoomUser{
				Uid:   req.Uid,
				Token: utils.UUID(),
			},
		}
	} else if req.GroupId != "" {
		//TODO ...
		//查找所有成员加入到房间
		//仅支持10人以下的群
	}

	var client protocol.RoomServiceClient
	s.RLock()
	for _, v := range s.roomClients {
		client = v
		break
	}
	s.RUnlock()

	resp, err := client.CreateRoom(ctx, &protocol.CreateRoomReq{
		Users: users,
	})

	ack.Addr = resp.Addr
	ack.RoomId = resp.RoomId
	ack.Token = req.Header.Token

	//推送给其它成员
	for _, u := range users {
		s.NotifyMessage(u.Uid, "RealTimeNotify", &protocol.RealTimeNotify{
			Header: &protocol.NotiHeader{},
			RealTimeInfo: &protocol.RealTimeInfo{
				Uid:     req.Uid,
				GroupId: req.GroupId,
				Token:   u.Token,
				RoomId:  resp.RoomId,
				Addr:    resp.Addr,
			},
			IsConnect: true,
		})
	}

	ack.Header.Code = 200
	ack.Header.Msg = "success"
	return
}

func (s *ChatLogic) CancelRealTime(req *protocol.CancelRealTimeReq, ack *protocol.CancelRealTimeAck) (err error) {
	//查找当前房间
	rid, addr, err := cache.GetUserRoomInfo(req.Header.Uid)
	if err != nil {
		return err
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, ok := s.roomClients[addr]
	if ok {
		client.LeaveRoom(ctx, &protocol.LeaveRoomReq{
			Uid:    req.Header.Uid,
			RoomId: rid,
		})
	}

	//删除信息
	cache.LeaveRoom(req.Header.Uid)
	uids, err := cache.GetRoomMembers(rid)
	if err != nil {
		log.Error(err)
	}
	if len(uids) <= 1 {
		//删除房间
		cache.DiscardRoom(rid)
	}
	//通知其他人
	for _, uid := range uids {
		s.NotifyMessage(uid, "RealTimeNotify", &protocol.RealTimeNotify{
			Header: &protocol.NotiHeader{},
			RealTimeInfo: &protocol.RealTimeInfo{
				RoomId: rid,
			},
			IsConnect: false,
		})
	}
	return nil
}

func (s *ChatLogic) NotifyMessage(uid, msgName string, msg interface{}) {
	var ok = false

	var err error
	if s.Notify != nil {
		ok, err = s.Notify(uid, msg)
		if err != nil {
			log.Errorf("notify msg error:%v", err)
		}
	}

	//没有长连接或发送失败 消息存到redis 前端轮询接收
	if !ok || err != nil {
		var body, err = json.Marshal(msg)
		if err != nil {
			return
		}
		var cacheMessage = protocol.CacheMessage{
			Name:      msgName,
			Body:      string(body),
			TimeStamp: time.Now().Unix(),
		}
		err = cache.CacheNotifyMessage(uid, cacheMessage)
		if err != nil {
			log.Errorf("cache notify msg error:%v", err)
		}
	}
}

func (s *ChatLogic) PollNotify(req *protocol.PollNotifyReq, ack *protocol.PollNotifyAck) error {
	msgs, err := cache.GetAllNotifyMessage(req.Header.Uid)
	if err != nil {
		return err
	}
	ack.Msg = msgs
	ack.Header.Code = 200
	ack.Header.Msg = "success"
	return err
}

func (s *ChatLogic) Watch() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("watcher paniced, recover:", err)
		}
	}()
	for {
		select {
		case <-s.stop:
			return
		default:
			res, err := s.watcher.Next()
			if err != nil {
				log.Errorf("watch error:%v", err)
				continue
			}

			var name = res.Service.Name

			if name == "live-chat.voip" {
				switch res.Action {
				case "create":
					for _, node := range res.Service.Nodes {
						s.AddRoomConn(node.Address)
					}
				case "delete":
					for _, node := range res.Service.Nodes {
						s.DelRoomConn(node.Address)
					}
				case "update":
					for _, node := range res.Service.Nodes {
						s.UpdateRoomConn(node.Address)
					}
				}
			}
		}
	}
}

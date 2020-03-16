package logic

import (
	"context"
	"encoding/json"
	"go-chat/cache"
	"go-chat/protocol"
	"go-lib/log"
	"go-lib/utils"
	"time"
)

//发送消息/邀请等。。。
type ChatLogic struct {
	roomClient protocol.RoomServiceClient
	Notify     func(uid string, msg interface{}) (bool, error)
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
		s.roomClient.CreateRoom(ctx,
			&protocol.CreateRoomReq{
				Users: users,
			})
	} else if req.GroupId != "" {
		//TODO ...
		//查找所有成员加入到房间
		//仅支持10人以下的群
	}

	resp, err := s.roomClient.CreateRoom(ctx, &protocol.CreateRoomReq{
		Users: users,
	})

	var addr = ""
	ack.Addr = addr
	ack.RoomId = resp.RoomId
	ack.Token = req.Header.Token

	//推送给其它成员
	for _, u := range users {
		s.NotifyMessage(u.Uid, "MessageNotify", &protocol.MessageNotify{
			Header:      &protocol.NotiHeader{},
			MessageType: int32(protocol.MessageType_RealTime),
			RealTimeInfo: &protocol.RealTimeInfo{
				Uid:     req.Uid,
				GroupId: req.GroupId,
				Token:   u.Token,
				RoomId:  resp.RoomId,
				Addr:    addr,
			},
		})
		// reals[u.Uid] =
	}

	ack.Header.Code = 200
	ack.Header.Msg = "success"
	return
}

func (s *ChatLogic) NotifyMessage(uid, msgName string, msg interface{}) {
	var ok = false

	ok, err := s.Notify(uid, msg)
	if err != nil {
		log.Errorf("notify msg error:%v", err)
	}

	//没有长连接或发送失败 消息存到redis 前端轮询接收
	if !ok {
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

package logic

import (
	"chat/cache"
	"chat/models"
	"chat/protocol"
	"context"
	"encoding/json"
	"go-lib/db"
	"go-lib/log"
	"go-lib/registry"
	"go-lib/utils"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type NotifyFunc func(uid int64, msg interface{}) (bool, error)

//发送消息/邀请等。。。
type ChatLogic struct {
	sync.RWMutex
	watcher     registry.Watcher
	Registry    registry.Registry
	roomClients map[string]protocol.RoomServiceClient
	Notify      NotifyFunc
	stop        chan bool
}

func NewChatLogic(regist registry.Registry, notifyFuncs ...NotifyFunc) *ChatLogic {
	var notifyFunc NotifyFunc = nil
	if len(notifyFuncs) > 0 {
		notifyFunc = notifyFuncs[0]
	}
	var watcher, err = regist.Watch()
	if err != nil {
		log.Error(err)
	}
	var c = &ChatLogic{
		roomClients: make(map[string]protocol.RoomServiceClient, 2),
		Notify:      notifyFunc,
		watcher:     watcher,
		Registry:    regist,
	}

	go c.Watch()
	return c
}

func (s *ChatLogic) SendMessage(r Reqer, a Acker) (err error) {
	req, _ := r.(*protocol.SendMessageReq)
	ack, _ := a.(*protocol.SendMessageAck)

	var uids = make([]int64, 1, 2)
	uids[0] = req.Header.Uid
	// var users = make(*models.User, 0, 1)
	//查找用户/群
	if req.ContactId != "" {
		_id, _ := primitive.ObjectIDFromHex(req.ContactId)
		var contact = &models.UserContact{
			Id: _id,
		}
		var ctx, cancel = NewContext()
		defer cancel()
		err = db.Mongo.Collection(contact.TableName(req.Header.Uid)).
			FindOne(ctx, bson.M{"_id": _id}).
			Decode(&contact)
		if err != nil {
			return Error(ack, err)
		}
		uids = append(uids, contact.UserId)
	}

	if req.GroupId != "" {
		_id, _ := primitive.ObjectIDFromHex(req.GroupId)
		var userGroup = &models.UserGroup{
			Id: _id,
		}
		var ctx, cancel = NewContext()
		defer cancel()
		err := db.Mongo.Collection(userGroup.TableName(req.Header.Uid)).FindOne(ctx, bson.M{"_id": _id}).Decode(&userGroup)
		if err != nil {
			return Error(ack, err)
		}
		//查找群成员
		var group = &models.Group{}
		err = db.Mongo.Collection(group.TableName()).FindOne(ctx, bson.M{"_id": _id}).Decode(&group)
		if err != nil {
			return Error(ack, err)
		}
		for _, v := range group.Members {
			uids = append(uids, v.UserId)
		}
	}

	s.NotifyMessage(uids, "msg", req.TextMessage)
	return Success(ack)
}

func (s *ChatLogic) RealTime(r Reqer, a Acker) (err error) {
	// switch req.Protocol {
	// case "tcp":
	// case "ws":
	// }
	req, _ := r.(*protocol.RealTimeReq)
	ack, _ := a.(*protocol.RealTimeAck)

	var users []*protocol.RoomUser

	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if req.Uid != 0 {
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

	ack.TcpAddr = resp.TcpAddr
	ack.WsAddr = resp.WsAddr
	ack.HttpAddr = resp.HttpAddr
	ack.RoomId = resp.RoomId
	ack.Token = req.Header.Token

	//推送给其它成员
	for _, u := range users {
		if u.Uid == req.Header.Uid {
			continue
		}
		s.NotifyMessage([]int64{u.Uid}, "RealTimeNotify", &protocol.RealTimeNotify{
			Header: &protocol.NotiHeader{},
			RealTimeInfo: &protocol.RealTimeInfo{
				Uid:      req.Uid,
				GroupId:  req.GroupId,
				Token:    u.Token,
				RoomId:   resp.RoomId,
				TcpAddr:  resp.TcpAddr,
				WsAddr:   resp.WsAddr,
				HttpAddr: resp.HttpAddr,
			},
			IsConnect: true,
		})
	}

	return Success(ack)
}

func (s *ChatLogic) CancelRealTime(r Reqer, a Acker) (err error) {
	req, _ := r.(*protocol.CancelRealTimeReq)
	ack, _ := a.(*protocol.CancelRealTimeAck)
	//查找当前房间
	rid, addr, err := cache.GetUserRoomInfo(req.Header.Uid)
	if err != nil {
		return Error(ack, err)
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
	s.NotifyMessage(uids, "RealTimeNotify", &protocol.RealTimeNotify{
		Header: &protocol.NotiHeader{},
		RealTimeInfo: &protocol.RealTimeInfo{
			RoomId: rid,
		},
		IsConnect: false,
	})
	ack.Header.Code = 200
	return Success(ack)
}

func (s *ChatLogic) NotifyMessage(uids []int64, msgName string, msg interface{}) {
	var ok = false
	var err error

	var cacheMessage []byte

	if s.Notify != nil {
		for _, uid := range uids {
			ok, err = s.Notify(uid, msg)
			if err != nil {
				log.Errorf("notify msg error:%v", err)
			}
			if !ok || err != nil {
				if cacheMessage == nil {
					if body, err := json.Marshal(msg); err != nil {
						log.Error(err)
					} else {
						cacheMessage, err = json.Marshal(protocol.CacheMessage{
							Name:      msgName,
							Body:      string(body),
							TimeStamp: time.Now().Unix(),
						})
						if err != nil {
							log.Error(err)
						}
					}
				}

				if len(cacheMessage) > 10 {
					err = cache.CacheNotifyMessage(uid, cacheMessage)
					if err != nil {
						log.Errorf("cache notify msg error:%v", err)
					}
				}
			}
		}

	}

	//没有长连接或发送失败 消息存到redis 前端轮询接收
	if !ok || err != nil {
		for _, uid := range uids {
			if cacheMessage == nil {
				if body, err := json.Marshal(msg); err != nil {
					log.Error(err)
				} else {
					cacheMessage, err = json.Marshal(protocol.CacheMessage{
						Name:      msgName,
						Body:      string(body),
						TimeStamp: time.Now().Unix(),
					})
					if err != nil {
						log.Error(err)
					}
				}
			}
			if len(cacheMessage) > 10 {
				err = cache.CacheNotifyMessage(uid, cacheMessage)
				if err != nil {
					log.Errorf("cache notify msg error:%v", err)
				}
			}
		}
	}
}

func (s *ChatLogic) CacheNotifyMessage(msg []byte, uid ...int64) {

}

func (s *ChatLogic) PollNotify(r Reqer, a Acker) error {
	req, _ := r.(*protocol.PollNotifyReq)
	ack, _ := a.(*protocol.PollNotifyAck)
	msgs, err := cache.GetAllNotifyMessage(req.Header.Uid)
	if err != nil {
		return Error(ack, err)
	}
	ack.Msg = msgs

	return Success(ack)
}

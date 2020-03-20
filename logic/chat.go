package logic

import (
	"chat/cache"
	"chat/logic/pool"
	"chat/models"
	"chat/protocol"
	"context"
	"go-lib/db"
	"go-lib/log"
	"go-lib/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type NotifyFunc func(uid int64, msg interface{}) (bool, error)

//发送消息/邀请等。。。
type ChatLogic struct {
	// sync.RWMutex
	// watcher     registry.Watcher
	// Registry    registry.Registry
	// roomClients map[string]protocol.RoomServiceClient
	Notify NotifyFunc
	stop   chan bool
}

func NewChatLogic(notifyFuncs ...NotifyFunc) *ChatLogic {
	var notifyFunc NotifyFunc = nil
	if len(notifyFuncs) > 0 {
		notifyFunc = notifyFuncs[0]
	}
	// var watcher, err = regist.Watch()
	// if err != nil {
	// 	log.Error(err)
	// }
	var c = &ChatLogic{
		// roomClients: make(map[string]protocol.RoomServiceClient, 2),
		Notify: notifyFunc,
		// watcher:     watcher,
		// Registry:    regist,
	}

	// go c.Watch()
	return c
}

func (s *ChatLogic) SendMessage(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.SendMessageReq)
	ack := &protocol.SendMessageAck{}
	var now = time.Now().Unix()

	// var users = make(*models.User, 0, 1)
	//查找用户/群
	if req.ContactId != "" {
		_id, _ := primitive.ObjectIDFromHex(req.ContactId)
		var contact = &models.UserContact{
			Id: _id,
		}
		var ctx, cancel = NewContext()
		defer cancel()
		err := db.Mongo.Collection(contact.TableName(req.Header.Uid)).
			FindOne(ctx, bson.M{"_id": _id}).
			Decode(&contact)

		//插入mongo
		var userMessage = &models.UserMessage{
			Text:        req.TextMessage,
			FromUid:     req.Header.Uid,
			ToUid:       contact.UserId,
			Read:        false,
			MessageType: int32(req.MessageType),
			CreateTime:  now,
		}
		ctx, cancel = NewContext()
		defer cancel()
		_, err = db.Mongo.Collection(userMessage.TableName(contact.UserId)).InsertOne(ctx, userMessage)
		if err != nil {
			return Error(ack, err)
		}
		ctx, cancel = NewContext()
		defer cancel()
		_, err = db.Mongo.Collection(userMessage.TableName(req.Header.Uid)).InsertOne(ctx, userMessage)
		if err != nil {
			return Error(ack, err)
		}

	}

	if req.GroupId != "" {
		_id, _ := primitive.ObjectIDFromHex(req.GroupId)
		var userGroup = &models.UserGroup{}
		var ctx, cancel = NewContext()
		defer cancel()
		err := db.Mongo.Collection(userGroup.TableName(req.Header.Uid)).FindOne(ctx, bson.M{"_id": _id}).Decode(&userGroup)
		if err != nil {
			return Error(ack, err)
		}
		if userGroup.GroupId == "" {
			return Fail(ack, "no such group")
		}
		//插入消息
		var groupMessage = &models.GroupMessage{
			GroupId:     userGroup.GroupId,
			Text:        req.TextMessage,
			FromUid:     req.Header.Uid,
			MessageType: int32(req.MessageType),
			Memtions:    req.Memtions,
			CreateTime:  now,
		}
		ctx, cancel = NewContext()
		defer cancel()
		_, err = db.Mongo.Collection(groupMessage.TableName(userGroup.GroupId)).InsertOne(ctx, groupMessage)
		if err != nil {
			return Error(ack, err)
		}

	}

	// s.NotifyMessage(uids, "msg", req.TextMessage)
	return Success(ack)
}

func (s *ChatLogic) RealTime(r Reqer) (Acker, error) {
	// switch req.Protocol {
	// case "tcp":
	// case "ws":
	// }
	req, _ := r.(*protocol.RealTimeReq)
	ack := &protocol.RealTimeAck{}

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

	client, ok := s.GetRandomRoomClient()
	if !ok {
		return Fail(ack, "")
	}

	resp, err := client.CreateRoom(ctx, &protocol.CreateRoomReq{
		Users: users,
	})
	if err != nil {
		return Error(ack, err)
	}

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
		s.NotifyRealTime([]int64{u.Uid}, &protocol.RealTimeNotify{
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

func (s *ChatLogic) CancelRealTime(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.CancelRealTimeReq)
	ack := &protocol.CancelRealTimeAck{}
	//查找当前房间
	rid, addr, err := cache.GetUserRoomInfo(req.Header.Uid)
	if err != nil {
		return Error(ack, err)
	}
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// var client =
	client, ok := s.GetRoomClient(addr)
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
	s.NotifyRealTime(uids, &protocol.RealTimeNotify{
		Header: &protocol.NotiHeader{},
		RealTimeInfo: &protocol.RealTimeInfo{
			RoomId: rid,
		},
		IsConnect: false,
	})
	ack.Header.Code = 200
	return Success(ack)
}

func (s *ChatLogic) NotifyRealTime(uids []int64, msg *protocol.RealTimeNotify) {
	for _, uid := range uids {
		err := cache.CacheRealTimeNotify(uid, msg)
		if err != nil {
			log.Errorf("cache notify msg error:%v", err)
		}
	}
}

func (s *ChatLogic) Poll(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.PollReq)
	ack := &protocol.PollAck{}
	msgs, err := cache.GetAllNotifyMessage(req.Header.Uid)
	if err != nil {
		return Error(ack, err)
	}
	ack.Msgs = msgs

	return Success(ack)
}

func (s *ChatLogic) PollMessage(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.PollMessageReq)
	ack := &protocol.PollMessageAck{}
	var userMessages = make([]models.UserMessage, 0, 10)
	var ctx, cancel = NewContext()
	defer cancel()
	c, err := db.Mongo.Collection((&models.UserMessage{}).TableName(req.Header.Uid)).Find(ctx, bson.M{"read": false}, options.Find().SetSort(bson.M{"_id": -1}))
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return Error(ack, err)
		}
	}
	defer c.Close(context.TODO())
	for c.TryNext(context.TODO()) {
		var msg models.UserMessage
		err := c.Decode(&msg)
		if err != nil {
			log.Error(err)
			continue
		}
		userMessages = append(userMessages, msg)
	}

	ack.Data = make([]*protocol.PollMessageAck_DataItem, 0, len(userMessages))
	for _, v := range userMessages {
		ack.Data = append(ack.Data, &protocol.PollMessageAck_DataItem{
			Text:        v.Text,
			FromUid:     v.FromUid,
			ToUid:       v.ToUid,
			MediaUrl:    v.MediaUrl,
			MessageType: v.MessageType,
			CreateTime:  v.CreateTime,
		})
	}
	return Success(ack)
}

func (s *ChatLogic) GetRoomClient(addr string) (protocol.RoomServiceClient, bool) {
	var conn, ok = pool.DefaultGrpcConnPool.GetConn(addr)
	if !ok {
		return nil, false
	}
	return protocol.NewRoomServiceClient(conn), true
}

func (s *ChatLogic) GetRandomRoomClient() (protocol.RoomServiceClient, bool) {
	var addr, conn = pool.DefaultGrpcConnPool.GetRandomConn()
	if addr != "" {
		return nil, false
	}
	return protocol.NewRoomServiceClient(conn), true
}

package logic

import (
	"chat/cache"
	"chat/models"
	"chat/protocol"
	"context"
	"go-lib/db"
	"go-lib/pool"
	"go-lib/utils"
	"time"

	log "github.com/sirupsen/logrus"

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
	var c = &ChatLogic{
		Notify: notifyFunc,
	}

	return c
}

func (s *ChatLogic) SendMessage(r Reqer) (Acker, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("SendMessage recovered", err)
		}
	}()
	req, _ := r.(*protocol.SendMessageReq)
	ack := &protocol.SendMessageAck{}
	var now = time.Now().Unix()

	var uids = make([]int64, 0, 1)
	var msg = &protocol.MessageNotify{
		Header: &protocol.NotifyHeader{},
		Body:   req.Body,
	}
	//查找用户
	if req.Body.ContactId != 0 {
		uids = append(uids, req.Header.Uid)
		var contact = models.UserContact{
			Id:  req.Body.ContactId,
			Uid: req.Header.Uid,
		}
		ok, err := db.Mysql.Get(&contact)
		if !ok {
			return Fail(ack, "no such contact")
		}
		if err != nil {
			return Error(ack, err)
		}
		uids = append(uids, contact.UserId)
		//插入mongo
		var userMessage = &models.UserMessage{
			Id:          primitive.NewObjectID(),
			Text:        req.Body.Text,
			FromUid:     req.Header.Uid,
			ToUid:       contact.UserId,
			Read:        false,
			MessageType: int32(req.Body.MessageType),
			CreateTime:  now,
		}
		ctx, cancel := NewContext()
		defer cancel()
		_, err = db.Mongo.Collection(userMessage.TableName(contact.UserId)).InsertOne(ctx, userMessage)
		if err != nil {
			return Error(ack, err)
		}
		ctx, cancel = NewContext()
		defer cancel()
		userMessage.Read = true
		_, err = db.Mongo.Collection(userMessage.TableName(req.Header.Uid)).InsertOne(ctx, userMessage)
		if err != nil {
			return Error(ack, err)
		}

	}

	if req.Body.GroupId != 0 {

		var userGroup = models.UserGroup{
			Uid:     req.Header.Uid,
			GroupId: req.Body.GroupId,
		}

		ok, err := db.Mysql.Get(&userGroup)
		if !ok || userGroup.GroupId != req.Body.GroupId {
			return Fail(ack, "you are not in this group, cannot send msg")
		}
		if err != nil {
			return Error(ack, err)
		}

		//插入消息
		var groupMessage = &models.GroupMessage{
			Id:          primitive.NewObjectID(),
			GroupId:     userGroup.GroupId,
			Text:        req.Body.Text,
			FromUid:     req.Header.Uid,
			MessageType: int32(req.Body.MessageType),
			Memtions:    req.Body.Memtions,
			CreateTime:  now,
		}
		var ctx, cancel = NewContext()
		defer cancel()
		_, err = db.Mongo.Collection(groupMessage.TableName(userGroup.GroupId)).InsertOne(ctx, groupMessage)
		if err != nil {
			return Error(ack, err)
		}

	}

	s.NotifyMessage(uids, msg)
	return Success(ack)
}

func (s *ChatLogic) RealTime(r Reqer) (Acker, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("RealTime recovered", err)
		}
	}()
	req, _ := r.(*protocol.RealTimeReq)
	ack := &protocol.RealTimeAck{}

	var users []*protocol.RoomUser

	if req.ContactId != 0 {

		//查找联系人
		userContact, ok, err := models.GetContactById(req.Header.Uid, req.ContactId)
		if err != nil {
			return Error(ack, err)
		}
		if !ok {
			return Fail(ack, "pls add the user as your contact")
		}
		users = []*protocol.RoomUser{
			&protocol.RoomUser{
				Uid:   req.Header.Uid,
				Token: utils.UUID(),
			},
			&protocol.RoomUser{
				Uid:   userContact.UserId,
				Token: utils.UUID(),
			},
		}
	} else if req.GroupId != 0 {
		//TODO ...
		//查找所有成员加入到房间
		//仅支持10人以下的群
	} else {
		return Fail(ack, "cannot chat with nobody")
	}

	client, ok := s.GetRandomRoomClient()
	if !ok {
		return Fail(ack, "live server not avliable")
	}

	ctx, cancel := NewContext()
	defer cancel()
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
			Header: &protocol.NotifyHeader{},
			RealTimeInfo: &protocol.RealTimeInfo{
				Uid:      req.Header.Uid,
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
	defer func() {
		if err := recover(); err != nil {
			log.Error("CancelRealTime recovered", err)
		}
	}()
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
		Header: &protocol.NotifyHeader{},
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
		msg.Header.Uid = 0
		if ok, err := s.Notify(uid, msg); err == nil && ok {
			if err != nil {
				log.Error(err)
			}
			continue
		}
		err := cache.CacheRealTimeNotify(uid, msg)
		if err != nil {
			log.Errorf("cache notify msg error:%v", err)
		}
	}
}

func (s *ChatLogic) NotifyMessage(uids []int64, msg Notifier) {
	for _, uid := range uids {
		header := msg.GetHeader()
		if header != nil {
			header.Uid = 0
		}

		ok, err := s.Notify(uid, msg)
		if err != nil {
			log.Error(err)
		}
		if !ok {
			log.Warn("notify fail, uid:%d", uid)
		}
	}

}

func (s *ChatLogic) Poll(r Reqer) (Acker, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("Poll recovered", err)
		}
	}()
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
	defer func() {
		if err := recover(); err != nil {
			log.Error("PollMessage recovered", err)
		}
	}()
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

	ack.Data = make([]*protocol.MessageAckBody, 0, len(userMessages))
	for _, v := range userMessages {
		ack.Data = append(ack.Data, &protocol.MessageAckBody{
			Text:        v.Text,
			FromUid:     v.FromUid,
			ToUid:       v.ToUid,
			MediaUrl:    v.MediaUrl,
			MessageType: v.MessageType,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
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
	if addr == "" {
		return nil, false
	}
	return protocol.NewRoomServiceClient(conn), true
}

func (s *ChatLogic) GetMessageUser(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.GetMessageUserReq)
	ack := &protocol.GetMessageUserAck{
		Header: &protocol.AckHeader{},
	}
	err := models.GetMessageUser(req.Header.Uid, ack)
	if err != nil {
		return Error(ack, err)
	}
	return Success(ack)
}

func (s *ChatLogic) GetUserMessage(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.GetMessageReq)
	ack := &protocol.GetMessageAck{
		Header: &protocol.AckHeader{},
	}
	err := models.GetUserMessage(req.Header.Uid, req.UserId, ack)
	if err != nil {
		return Error(ack, err)
	}
	ack.UserId = req.UserId
	ack.GroupId = req.GroupId
	return Success(ack)
}

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
)

type NotifyFunc func(uid int64, msg interface{}) (bool, error)

//发送消息/邀请等。。。
type ChatLogic struct {
	// sync.RWMutex
	// watcher     registry.Watcher
	// Registry    registry.Registry
	// roomClients map[string]protocol.RoomServiceClient
	Notifys []NotifyFunc
	stop    chan bool
}

func NewChatLogic(notifyFuncs ...NotifyFunc) *ChatLogic {
	// var notifyFunc NotifyFunc = nil
	// if len(notifyFuncs) > 0 {
	// 	notifyFunc = notifyFuncs
	// }
	var c = &ChatLogic{
		Notifys: make([]NotifyFunc, 0, 2),
	}
	if notifyFuncs != nil {
		for _, v := range notifyFuncs {
			if v != nil {
				c.Notifys = append(c.Notifys, v)
			}
		}
	}

	return c
}

func (s *ChatLogic) SendMessage(r protocol.Reqer) (protocol.Acker, error) {
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
	msg.FromUserId = req.Header.Uid
	//查找用户
	if req.Body.ToUserId != 0 {
		uids = append(uids, req.Header.Uid)
		var contact = models.UserContact{
			UserId: req.Body.ToUserId,
			Uid:    req.Header.Uid,
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

	if req.Body.ToGroupId != 0 {

		var userGroup = models.UserGroup{
			Uid:     req.Header.Uid,
			GroupId: req.Body.ToGroupId,
		}

		ok, err := db.Mysql.Get(&userGroup)
		if !ok || userGroup.GroupId != req.Body.ToGroupId {
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

	//查找发送者信息
	user, _ := models.FindOneUserByUid(req.Header.Uid)
	msg.Nickname = user.Nickname
	msg.Avatar = user.Avatart
	s.NotifyMessage(uids, msg)
	return Success(ack)
}

func (s *ChatLogic) RealTime(r protocol.Reqer) (protocol.Acker, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("RealTime recovered", err)
		}
	}()
	req, _ := r.(*protocol.RealTimeReq)
	ack := &protocol.RealTimeAck{}

	var users []*protocol.RoomUser

	if req.UserId != 0 {

		//查找联系人
		userContact, ok, err := models.GetContactByUserId(req.Header.Uid, req.UserId)
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

	//推送给其它成员
	for _, u := range users {
		if u.Uid == req.Header.Uid {
			ack.Token = u.Token
			continue
		}
		s.NotifyRealTime([]int64{u.Uid}, &protocol.RealTimeNotify{
			Header: &protocol.NotifyHeader{},
			RealTimeInfo: &protocol.RealTimeInfo{
				UserId:   req.Header.Uid,
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

func (s *ChatLogic) CancelRealTime(r protocol.Reqer) (protocol.Acker, error) {
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
		var notified bool
		for _, notify := range s.Notifys {
			ok, err := notify(uid, msg)
			if err != nil {
				log.Error(err)
			}
			if ok {
				notified = true
			}
		}
		if !notified {
			err := cache.CacheRealTimeNotify(uid, msg)
			if err != nil {
				log.Errorf("cache notify msg error:%v", err)
			}
		}

	}
}

func (s *ChatLogic) NotifyMessage(uids []int64, msg protocol.Notifier) {
	for _, uid := range uids {
		header := msg.GetHeader()
		if header != nil {
			header.Uid = 0
		}

		for _, notify := range s.Notifys {
			_, err := notify(uid, msg)
			if err != nil {
				log.Error(err)
			}
		}
	}

}

func (s *ChatLogic) Poll(r protocol.Reqer) (protocol.Acker, error) {
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

func (s *ChatLogic) GetMessageUser(r protocol.Reqer) (protocol.Acker, error) {
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

func (s *ChatLogic) GetUserMessage(r protocol.Reqer) (protocol.Acker, error) {
	req, _ := r.(*protocol.GetMessageReq)
	ack := &protocol.GetMessageAck{
		Header: &protocol.AckHeader{},
	}
	err := models.GetUserMessage(req.Header.Uid, req.UserId, ack)
	if err != nil {
		return Error(ack, err)
	}
	return Success(ack)
}

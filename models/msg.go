package models

import (
	"chat/protocol"
	"context"
	"fmt"
	"go-lib/db"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

const (
	BodyTypeText int32 = 1 + iota
	BodyTypePicture
	BodyTypeAudio
	BodyTypeVideo
	BodyTypeLink
)

type SystemMessage struct {
	Id          primitive.ObjectID `bson:"_id"`
	Text        string
	MediaUrl    string
	BodyType    int32
	MessageType int32
	Link        LinkData
	ToUids      []int64
	CreateTime  int64
}

type LinkData struct {
	Url    string //跳转链接
	Title  string //标题
	Body   string //预览
	Avatar string //缩略图
}

func (m *SystemMessage) TableName(kind int32) string {
	return fmt.Sprintf("system_message:%d", kind)
}

type UserMessage struct {
	Id          primitive.ObjectID `bson:"_id"`
	Text        string
	MediaUrl    string
	MessageType int32
	Link        LinkData
	FromUid     int64
	ToUid       int64
	GroupId     string
	Read        bool
	Deleted     bool

	CreateTime int64
	UpdateTime int64
}

func (m *UserMessage) TableName(uid int64) string {
	return fmt.Sprintf("user_message_%d", uid)
}

type GroupMessage struct {
	Id          primitive.ObjectID `bson:"_id"`
	Text        string
	MediaUrl    string
	MessageType int32
	Link        LinkData
	FromUid     int64
	GroupId     int64
	Memtions    []*protocol.Memtion
	Deleted     bool

	CreateTime int64
	UpdateTime int64
}

func (m *GroupMessage) TableName(groupId int64) string {
	return fmt.Sprintf("group_message_%s", groupId)
}

func NewContext(ttls ...time.Duration) (context.Context, context.CancelFunc) {
	var ttl = 5 * time.Second
	if len(ttls) > 0 {
		ttl = ttls[0]
	}
	return context.WithTimeout(context.Background(), ttl)
}

func GetMessageUser(uid int64, ack *protocol.GetMessageUserAck) error {
	if ack.Users == nil {
		ack.Users = make([]*protocol.GetMessageUserAck_MessageUser, 0, 1)
	}
	//查找未读消息
	var userMessage = &UserMessage{}

	var ctx, cancel = NewContext()
	defer cancel()
	var tbName = userMessage.TableName(uid)
	rows, err := db.Mongo.Collection(tbName).Find(ctx, bson.M{"read": false})
	if err != nil {
		return err
	}
	var m = make(map[int64][]UserMessage)
	for rows.Next(context.TODO()) {
		var msg UserMessage
		err := rows.Decode(&msg)
		if err != nil {
			continue
		}

		if msg.FromUid == uid { //发出去的消息
			if _, ok := m[msg.ToUid]; ok {
				m[msg.ToUid] = append(m[msg.ToUid], msg)
			} else {
				m[msg.ToUid] = []UserMessage{msg}
			}
		} else {
			if _, ok := m[msg.FromUid]; ok {
				m[msg.FromUid] = append(m[msg.FromUid], msg)
			} else {
				m[msg.FromUid] = []UserMessage{msg}
			}
		}
	}
	//分组
	var uids = make([]int64, 0, len(m))
	for k, v := range m {
		uids = append(uids, k)
		var msgs = make([]*protocol.MessageAckBody, 0, len(v))
		for _, msg := range v {
			msgs = append(msgs, &protocol.MessageAckBody{
				Text:       msg.Text,
				FromUid:    msg.FromUid,
				ToUid:      msg.ToUid,
				CreateTime: msg.CreateTime,
				UpdateTime: msg.UpdateTime,
			})
		}
		ack.Users = append(ack.Users, &protocol.GetMessageUserAck_MessageUser{
			UserId:   k,
			Count:    int32(len(v)),
			Messages: msgs,
		})
	}
	//查找用户信息
	users, err := FindUsersByUids(uids)
	if err != nil {
		log.Error(err)
		return nil
	}
	for _, v := range ack.Users {
		var user = users[v.UserId]
		v.Nickname = user.Nickname
		v.Avatar = user.Avatart
	}
	return nil
}

func GetUserMessage(uid, toUid int64, ack *protocol.GetMessageAck) error {
	if ack.Messages == nil {
		ack.Messages = make([]*protocol.MessageAckBody, 0, 10)
	}
	//查找未读消息
	var userMessage = &UserMessage{}

	var ctx, cancel = NewContext()
	defer cancel()
	var tbName = userMessage.TableName(uid)
	rows, err := db.Mongo.Collection(tbName).Find(ctx, bson.M{"$or": []bson.M{
		bson.M{"fromuid": toUid},
		bson.M{"touid": toUid},
	}})
	if err != nil {
		return err
	}

	for rows.Next(context.TODO()) {
		var msg = new(UserMessage)
		err := rows.Decode(msg)
		if err != nil {
			log.Error(err)
			continue
		}
		ack.Messages = append(ack.Messages, &protocol.MessageAckBody{
			Text:       msg.Text,
			FromUid:    msg.FromUid,
			ToUid:      msg.ToUid,
			CreateTime: msg.CreateTime,
			UpdateTime: msg.UpdateTime,
		})
	}
	return nil
}

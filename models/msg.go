package models

import (
	"chat/protocol"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

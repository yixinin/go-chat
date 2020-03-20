package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	GroupTypeTemp   uint8 = 1 //限时
	GroupTypeNormal uint8 = 2 //永久
)

type Group struct {
	Id     primitive.ObjectID `bson:"_id"`
	Name   string
	Avatar string

	GroupType uint8 //群类型

	Members []*GroupMember

	CreateBy int64 //创建人

	Log        string //群公告
	ExpireTime int64

	CreateTime int64
	UpdateTime int64
}

func (g *Group) TableName() string {
	return TablePrefix + "_group"
}

type GroupMember struct {
	UserId   int64
	Remarks  string //群昵称
	IsAdmin  bool
	JoinDesc JoinDesc //如何加入
	JoinTime time.Time
}

type JoinDesc struct {
	InviteUserId  string //邀请人
	InviteUrl     string //邀请链接
	ApproveUserId string //审核人
}

//
type GroupAuth struct {
	Id           primitive.ObjectID `bson:"_id"`
	UserId       int64
	Pass         bool
	PassAdmin    int64 //通过人
	InviteCode   string
	InviteUserId int64
	Token        string
}

func (g GroupAuth) TableName() string {
	return TablePrefix + "_group_auth"
}

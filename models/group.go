package models

import (
	"fmt"
	"time"
)

const (
	GroupTypeNormal uint8 = 1 //永久
)

type Group struct {
	Id     int64  `xorm:"pk autoincr"`
	Name   string `xorm:"unique"`
	Avatar string

	GroupType uint8 `xorm:"default(1)"` //群类型

	CreateBy int64 //创建人

	Log  string //群公告
	Logs string //历史公告

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

func (g *Group) TableName() string {
	return "group"
}

//
type GroupMember struct {
	GroupId       int64  `xorm:"-"`
	Id            int64  `xorm:"pk autoincr"`
	UserId        int64  `xorm:"unique"`
	GroupNickname string //群昵称

	Approved      bool
	ApproveUserId string //审核人

	InviteCode   string
	InviteUserId string //邀请人
	InviteUrl    string //邀请链接

}

func (m *GroupMember) TableName() string {
	return fmt.Sprintf("group_auth_%d", m.GroupId)
}

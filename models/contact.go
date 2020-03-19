package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	NotifyLevelMute    uint8 = 0 //静音 不提示消息
	NotifyLevelDisplay uint8 = 1 //仅显示， 不提示
	NotifyLevelNormal  uint8 = 2 //打开提示消息
	NotifyLevelPush    uint8 = 3 //推送消息
)

type Contact struct {
	Id         primitive.ObjectID `bson:"_id"`
	UserAId    int64              //x用户的联系人列表
	UserBId    int64
	RemarksA   string //预设备注备注
	Status     int32  //状态 1=待通过 2=已拒绝 3=已通过 4=已过期
	CreateTime int64
	UpdateTime int64
}

func (Contact) TableName() string {
	return "chat_contact"
}

//集合名：user_contact_userId
type UserContact struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      int64              //联系人
	Sort        int
	IsFavorites bool
	Remarks     string //备注
	NotifyLevel uint8

	CreateTime int64
	UpdateTime int64
}

func (UserContact) TableName(uid int64) string {
	return fmt.Sprintf("chat_user_contact_%d", uid)
}

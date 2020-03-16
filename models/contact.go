package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	NotifyLevelMute    uint8 = 0 //静音 不提示消息
	NotifyLevelDisplay uint8 = 1 //仅显示， 不提示
	NotifyLevelNormal  uint8 = 2 //打开提示消息
	NotifyLevelPush    uint8 = 3 //推送消息
)

type Contact struct {
	Id            primitive.ObjectID `bson:"_id"`
	UserId        string             //x用户的联系人列表
	ContactUserId string             //x的联系人
	Sort          int
	IsFavorites   bool
	Remarks       string //备注
	NotifyLevel   uint8

	CreateTime int64
	UpdateTime int64
}

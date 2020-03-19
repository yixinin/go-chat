package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// user_group_userId
type UserGroup struct {
	Id          primitive.ObjectID `bson:"_id"`
	GroupId     string
	Sort        int
	IsFavorites bool
	Remarks     string //备注
	NotifyLevel int8

	CreateTime int64
	UpdateTime int64
}

func (g *UserGroup) TableName(uid int64) string {
	return fmt.Sprintf("chat_user_group_%d", uid)
}

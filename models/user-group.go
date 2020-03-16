package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserGroup struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      string
	GroupId     string
	Sort        int
	IsFavorites bool
	Remarks     string //备注
	NotifyLevel int8

	CreateTime int64
	UpdateTime int64
}

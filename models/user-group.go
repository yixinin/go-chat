package models

import "time"

type UserGroup struct {
	Id         int64
	UserId     int64
	GroupId    int64
	IsAdmin    bool
	JoinDesc   string //如何加入
	CreateTime time.Time
	UpdateTime time.Time
}

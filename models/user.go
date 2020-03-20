package models

import (
	"go-lib/db"
	"time"
)

type User struct {
	Id           int64 `xorm:"pk autoincr"`
	Username     string
	PasswordHash string
	DevideCode   string
	Nickname     string
	Avatart      string

	InviteCode string //可以用来邀请其他人

	DeleteTime int64 //删除账户时间

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

func SyncTables() {
	db.Mysql.Sync2(new(User))
}

package models

import (
	"fmt"
	"go-lib/db"
	"go-lib/log"
	"time"
)

const (
	TablePrefix = "chat"
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

func (*User) TableName() string {
	return fmt.Sprintf("%s_user", TablePrefix)
}

func SyncTables() {
	err := db.Mysql.Sync2(new(User))
	if err != nil {
		log.Errorf("sync tables error:%v", err)
	}
}

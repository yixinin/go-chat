package models

import (
	"fmt"
	"time"
)

// user_group_userId
type UserGroup struct {
	Uid         int64 `xorm:"-"`
	Id          int64 `xorm:"pk autoincr"`
	GroupId     int64 `xorm:"unique"`
	Sort        int
	IsFavorites bool
	NotifyLevel int8

	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

func (g *UserGroup) TableName() string {
	if g.Uid == 0 {
		return ""
	}
	return fmt.Sprintf("user_group_%d", g.Uid)
}

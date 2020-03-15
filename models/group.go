package models

import (
	"time"
)

type Group struct {
	Id   int64
	Name string

	CreateBy int64 //创建人

	CreateTime time.Time
	UpdateTime time.Time
}

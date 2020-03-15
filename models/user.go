package models

import (
	"time"
)

type User struct {
	Id           int64
	Username     string
	PasswordHash string
	DevideCode   string
	Nickname     string
	Avatart      string

	CreateTime time.Time
	UpdateTime time.Time
}

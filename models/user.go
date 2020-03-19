package models

type User struct {
	Id           int64
	Username     string
	PasswordHash string
	DevideCode   string
	Nickname     string
	Avatart      string

	InviteCode string //可以用来邀请其他人

	DeleteTime int64 //删除账户时间

	CreateTime int64
	UpdateTime int64
}

package models

import (
	"chat/cache"
	"chat/protocol"
	"go-lib/db"
	"go-lib/utils"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-xorm/xorm"
)

const (
	DefaultAvatar = "http://localhost:8080/static/avatar/default.jpg"
)

type User struct {
	Id           int64  `xorm:"pk autoincr"`
	Username     string `xorm:"unique"`
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
	return "user"
}

func CreateUser(req *protocol.SignUpReq, ack *protocol.SignUpAck) (msg string, err error) {
	var now = time.Now()
	if req.DeviceCode == "" {
		req.DeviceCode = utils.UUID()
	}

	var sess = db.Mysql.NewSession()
	defer sess.Close()
	err = sess.Begin()
	if err != nil {
		return
	}
	defer sess.Rollback()

	uid, err := sess.InsertOne(&User{
		Username:     req.Username,
		PasswordHash: utils.MD5(req.Password),
		DevideCode:   req.DeviceCode,
		Nickname:     req.Nickname,
		Avatart:      DefaultAvatar,
		InviteCode:   utils.UUID(),
		CreateTime:   now,
		UpdateTime:   now,
	})
	if err != nil {
		log.Error(err)
		return "username is already taken", nil
	}
	if uid == 0 {
		return "siginup failed, pls try later", nil
	}
	//创建表
	err = createUserTables(sess, uid)
	if err != nil {
		return "", err
	}

	//设置登录缓存
	token, _, err1 := cache.SetToken(uid, req.DeviceType)
	if err1 != nil {
		log.Error("set user token error:%v", err)
	}
	ack.Token = token
	ack.Header.Uid = uid
	return "", nil
}

func createUserTables(sess *xorm.Session, uid int64) error {
	err := sess.CreateTable(&UserContact{Uid: uid})
	if err != nil {
		return err
	}
	return sess.CreateTable(&UserGroup{Uid: uid})
}

package models

import (
	"chat/cache"
	"chat/protocol"
	"go-lib/db"
	"go-lib/utils"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-xorm/xorm"
)

var avatars = []string{
	"http://localhost:8080/static/avatar/default0.jpg",
	"http://localhost:8080/static/avatar/default1.jpg",
	"http://localhost:8080/static/avatar/default2.jpg",
	"http://localhost:8080/static/avatar/default3.jpg",
	"http://localhost:8080/static/avatar/default4.jpg",
	"http://localhost:8080/static/avatar/default5.jpg",
	"http://localhost:8080/static/avatar/default6.jpg",
	"http://localhost:8080/static/avatar/default7.jpg",
	"http://localhost:8080/static/avatar/default8.jpg",
	"http://localhost:8080/static/avatar/default9.jpg",
}

func init() {
	rand.Seed(time.Now().Unix())
}

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

	var user = &User{
		Username:     req.Username,
		PasswordHash: utils.MD5(req.Password),
		DevideCode:   req.DeviceCode,
		Nickname:     req.Nickname,
		Avatart:      avatars[rand.Intn(10)],
		InviteCode:   utils.UUID(),
		CreateTime:   now,
		UpdateTime:   now,
	}
	_, err = sess.InsertOne(user)
	if err != nil {
		log.Error(err)
		return "username is already taken", nil
	}
	var uid = user.Id
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
	ack.UserInfo = &protocol.SignUserInfo{
		Avatar:   user.Avatart,
		Uid:      uid,
		Nickname: user.Nickname,
	}
	return "", nil
}

func createUserTables(sess *xorm.Session, uid int64) error {
	var uc = &UserContact{Uid: uid}
	err := sess.CreateTable(uc)
	if err != nil {
		return err
	}
	err = sess.CreateUniques(uc)
	if err != nil {
		return err
	}
	var ug = &UserGroup{Uid: uid}
	err = sess.CreateTable(ug)
	if err != nil {
		return err
	}
	err = sess.CreateUniques(ug)
	return err
}

func FindUsersByUids(uids []int64) (map[int64]*User, error) {
	var users = make([]*User, 0, len(uids))
	err := db.Mysql.In("id", uids).Find(&users)
	if err != nil {
		return nil, err
	}
	var m = make(map[int64]*User, len(uids))
	for _, v := range users {
		m[v.Id] = v
	}
	return m, nil
}

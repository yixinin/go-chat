package logic

import (
	"chat/cache"
	"chat/models"
	"chat/protocol"
	"fmt"
	"go-lib/db"
	"go-lib/log"
	"go-lib/utils"
	"time"
)

const (
	DefaultAvatar = "livechat/avatar/default.png"
)

type AccountLogic struct {
}

//账号操作等。。。

//SignUp 注册
func (s *AccountLogic) SignUp(r Reqer, a Acker) (err error) {
	req, _ := r.(*protocol.SignUpReq)
	ack, _ := a.(*protocol.SignUpAck)

	var now = time.Now().Unix()
	if req.DeviceCode == "" {
		req.DeviceCode = utils.UUID()
	}

	var regType = 1
	var userName = req.Username
	var nickname = req.Username
	if req.Username == "" {
		regType = 2
		var device = req.DeviceCode
		if len(req.DeviceCode) > 10 {
			device = req.DeviceCode[:10]
		}
		nickname = fmt.Sprintf("游客(%s)", device)
		req.Username = req.DeviceCode
	}

	uid, err := db.Mysql.InsertOne(&models.User{
		Username:     userName,
		PasswordHash: utils.MD5(req.Password),
		DevideCode:   req.DeviceCode,
		Nickname:     nickname,
		Avatart:      DefaultAvatar,
		InviteCode:   utils.UUID(),
		CreateTime:   now,
		UpdateTime:   now,
	})
	if err != nil {
		if regType == 1 {
			ack.Header.Code = 400
			ack.Header.Msg = "repeate username"
		} else {
			ack.Header.Code = 400
			ack.Header.Msg = "repeate device"
		}
		return Error(ack, err)
	}

	token, _, err1 := cache.SetToken(uid, req.DeviceType)
	if err1 != nil {
		log.Error("set user token error:%v", err)
	}
	ack.Token = token

	return Success(ack)
}

//登录
func (s *AccountLogic) SignIn(r Reqer, a Acker) (err error) {
	req, _ := r.(*protocol.SignInReq)
	ack, _ := a.(*protocol.SignInAck)

	var user = models.User{
		Username: req.Username,
	}
	ok, err := db.Mysql.Get(&user)
	if !ok {

		return Fail(ack, "username and password not match")
	}
	if err != nil {
		return Error(ack, err)
	}

	if user.PasswordHash != utils.MD5(req.Password) {

		return Fail(ack, "username and password not match")
	}

	token, oldToken, err := cache.SetToken(user.Id, req.DeviceType)
	if err != nil {
		return Error(ack, err)
	}
	ack.Token = token
	if oldToken != "" {
		s.tryKickDevice(user.Id, req.DeviceType, token)
	}

	return Success(ack)
}

func (s *AccountLogic) SignOut(r Reqer, a Acker) (err error) {
	req, _ := r.(*protocol.SignOutReq)
	ack, _ := a.(*protocol.SignOutAck)
	var uid = req.Header.Uid
	deviceType, err := cache.GetDeviceToken(uid, req.Header.Token)
	cache.DelDevice(uid, deviceType)
	s.tryKickDevice(uid, deviceType, req.Header.Token)

	return Success(ack)
}

func (s *AccountLogic) Delete(r Reqer, a Acker) (err error) {
	req, _ := r.(*protocol.DeleteReq)
	ack, _ := a.(*protocol.DeleteAck)

	//标记删除
	var user = &models.User{
		Id:         req.Header.Uid,
		DeleteTime: time.Now().AddDate(0, 0, 7).Unix(),
	}
	n, err := db.Mysql.Cols("delete_time").Update(&user)

	if n != 1 {
		return Fail(ack, "operate fail, pls try later")
	}

	return Success(ack)
}

func (s *AccountLogic) ChangePassword(r Reqer, a Acker) (err error) {
	// req, _ := r.(*protocol.ChangePasswordReq)
	ack, _ := a.(*protocol.ChangePasswordAck)
	return Success(ack)
}

func (s *AccountLogic) ResetPassword(r Reqer, a Acker) (err error) {
	// req,_ :=r.(*protocol.ResetPasswordReq)
	ack, _ := a.(*protocol.ResetPasswordAck)
	return Success(ack)
}

func (s *AccountLogic) tryKickDevice(uid int64, deviceType int32, token string) {
	cache.DelToken(token)
	//断开长连接 断开通话

}

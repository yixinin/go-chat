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
	// acceptFunc AcceptFunc
}

func NewAccountLogic() *AccountLogic {
	return &AccountLogic{
		// acceptFunc: acc,
	}
}

//账号操作等。。。

//SignUp 注册
func (s *AccountLogic) SignUp(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.SignUpReq)
	ack := &protocol.SignUpAck{
		Header: &protocol.AckHeader{},
	}

	var now = time.Now()
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
	if len(req.Password) < 6 {
		return Fail(ack, "password si too simple")
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
	ack.Header.Uid = uid
	// s.acceptFunc(uid, nil)
	return Success(ack)
}

//登录
func (s *AccountLogic) SignIn(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.SignInReq)
	ack := &protocol.SignInAck{
		Header: &protocol.AckHeader{},
	}

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
	ack.Header.Uid = user.Id
	if oldToken != "" {
		s.tryKickDevice(user.Id, req.DeviceType, oldToken)
	}
	return Success(ack)
}

func (s *AccountLogic) SignOut(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.SignOutReq)
	ack := &protocol.SignOutAck{}
	var uid = req.Header.Uid
	deviceType, err := cache.GetDeviceToken(uid, req.Header.Token)
	if err != nil {
		return Error(ack, err)
	}
	cache.DelDevice(uid, deviceType)
	s.tryKickDevice(uid, deviceType, req.Header.Token)

	return Success(ack)
}

func (s *AccountLogic) Delete(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.DeleteReq)
	ack := &protocol.DeleteAck{}

	//标记删除
	var user = &models.User{
		Id:         req.Header.Uid,
		DeleteTime: time.Now().AddDate(0, 0, 7).Unix(),
	}
	n, err := db.Mysql.Cols("delete_time").Update(&user)
	if err != nil {
		return Error(ack, err)
	}
	if n != 1 {
		return Fail(ack, "operate fail, pls try later")
	}

	return Success(ack)
}

func (s *AccountLogic) ChangePassword(r Reqer) (ack Acker, err error) {
	// req, _ := r.(*protocol.ChangePasswordReq)
	ack = &protocol.ChangePasswordAck{}
	return Success(ack)
}

func (s *AccountLogic) ResetPassword(r Reqer) (ack Acker, err error) {
	// req,_ :=r.(*protocol.ResetPasswordReq)
	ack = &protocol.ResetPasswordAck{}
	return Success(ack)
}

func (s *AccountLogic) tryKickDevice(uid int64, deviceType int32, token string) {
	cache.DelToken(token)
	//断开长连接 断开通话

}

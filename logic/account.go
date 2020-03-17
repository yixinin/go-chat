package logic

import (
	"chat/cache"
	"chat/models"
	"chat/protocol"
	"context"
	"fmt"
	"go-lib/db"
	"go-lib/log"
	"go-lib/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const (
	DefaultAvatar = "livechat/avatar/default.png"
)

type AccountLogic struct {
}

//账号操作等。。。

//SignUp 注册
func (s *AccountLogic) SignUp(req *protocol.SignUpReq, ack *protocol.SignUpAck) (err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var now = time.Now().Unix()

	if req.DeviceCode == "" {
		req.DeviceCode = utils.UUID()
	}

	var regType = 1
	var userName = req.Username
	var nickname = req.Username
	if req.Username == "" {
		regType = 2
		nickname = fmt.Sprintf("游客(%s)", req.DeviceCode)
		req.Username = req.DeviceCode
	}

	ret, err := db.Mongo.Collection("user").InsertOne(ctx, &models.User{
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
		if err == mongo.ErrMultipleIndexDrop {
			if regType == 1 {
				ack.Header.Code = 400
				ack.Header.Msg = "repeate username"
			} else {
				ack.Header.Code = 400
				ack.Header.Msg = "repeate device"
			}

		}
		return err
	}
	uid, ok := ret.InsertedID.(primitive.ObjectID)
	if !ok {
		return fmt.Errorf("signup fail, pls try later")
	}

	token, _, err1 := cache.SetToken(uid.Hex(), req.DeviceType)
	if err1 != nil {
		log.Error("set user token error:%v", err)
	}
	ack.Token = token
	ack.Header.Code = 200
	ack.Header.Msg = "success"
	return
}

//登录
func (s *AccountLogic) SignIn(req *protocol.SignInReq, ack *protocol.SignInAck) (err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user models.User
	ex := db.Mongo.Collection("user").FindOne(ctx, bson.M{"username": req.Username}).Decode(&user)
	if ex != nil {
		if ex != mongo.ErrNoDocuments {
			ack.Header.Code = 401
			ack.Header.Msg = "username and password not match"
			return nil
		}
		return ex
	}

	if user.PasswordHash != utils.MD5(req.Password) {
		ack.Header.Code = 401
		ack.Header.Msg = "username and password not match"
		return
	}

	token, oldToken, err := cache.SetToken(user.Id.Hex(), req.DeviceType)
	if err != nil {
		return err
	}
	ack.Token = token
	if oldToken != "" {
		s.tryKickDevice(user.Id.Hex(), req.DeviceType, token)
	}
	ack.Header.Code = 200
	ack.Header.Msg = "success"
	return
}

func (s *AccountLogic) SignOut(req *protocol.SignOutReq, ack *protocol.SignOutAck) (err error) {
	var uid = req.Header.Uid
	deviceType, err := cache.GetDeviceToken(uid, req.Header.Token)
	cache.DelDevice(uid, deviceType)
	s.tryKickDevice(uid, deviceType, req.Header.Token)
	ack.Header.Code = 200
	ack.Header.Msg = "success"
	return
}

func (s *AccountLogic) Delete(req *protocol.DeleteReq, ack *protocol.DeleteAck) (err error) {
	uid, err := cache.GetToken(req.Header.Token)
	//标记删除
	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var id, _ = primitive.ObjectIDFromHex(uid)
	ret, err := db.Mongo.Collection("user").UpdateOne(ctx,
		bson.M{"_id": id},
		bson.M{"wait_delete": time.Now().AddDate(0, 0, 7).Unix()})
	if err != nil {
		return err
	}
	if ret.ModifiedCount <= 0 {
		ack.Header.Code = 400
		ack.Header.Msg = "operate fail, pls try later"
		return
	}
	ack.Header.Code = 200
	ack.Header.Msg = "success"
	return
}

func (s *AccountLogic) ChangePassword(req *protocol.ChangePasswordReq, ack *protocol.ChangePasswordAck) (err error) {
	ack.Header.Code = 200
	ack.Header.Msg = "success"
	return
}

func (s *AccountLogic) ResetPassword(req *protocol.ResetPasswordReq, ack *protocol.ResetPasswordAck) (err error) {
	ack.Header.Code = 200
	ack.Header.Msg = "success"
	return
}

func (s *AccountLogic) tryKickDevice(uid string, deviceType int32, token string) {
	cache.DelToken(token)
	//断开长连接 断开通话
}

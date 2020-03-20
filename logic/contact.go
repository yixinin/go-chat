package logic

import (
	"chat/models"
	"chat/protocol"
	"context"
	"fmt"
	"go-lib/db"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type ContactLogic struct {
}

func NewContactLogic() *ContactLogic {
	return &ContactLogic{}
}

func (s *ContactLogic) SearchUser(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.SearchUserReq)
	ack := &protocol.SearchUserAck{}

	var users []*models.User
	var where = fmt.Sprintf("username like %%%s%% or nickname like %%%s%%", req.Key, req.Key)
	err := db.Mysql.
		Where(where).
		Limit(10).
		Find(&users)
	if err != nil {
		return Error(ack, err)
	}

	var datas = make([]*protocol.SearchUserAck_DataItem, 0, 10)
	for _, user := range users {
		datas = append(datas, &protocol.SearchUserAck_DataItem{
			Username: user.Username,
			Nickname: user.Nickname,
			Avatar:   user.Avatart,
		})
	}
	ack.Data = datas
	return Success(ack)
}

func (s *ContactLogic) AddContact(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.AddContactReq)
	ack := &protocol.AddContactAck{}
	var now = time.Now().Unix()

	if req.AuthId != "" { //通过验证
		ctx, cancel := NewContext()
		defer cancel()
		var contact = &models.Contact{}
		_id, _ := primitive.ObjectIDFromHex(req.AuthId)
		err := db.Mongo.Collection(contact.TableName()).FindOne(ctx, bson.M{"_id": _id}).Decode(contact)
		if err != nil {
			return Error(ack, err)
		}
		if contact.UserBId != req.Header.Uid {
			return Fail(ack, "accsess deined")
		}
		//添加
		ctx, cancel = NewContext()
		defer cancel()
		userContact := &models.UserContact{
			UserId:     req.Header.Uid,
			Remarks:    req.SetRemarks,
			CreateTime: now,
			UpdateTime: now,
		}
		_, err = db.Mongo.Collection(userContact.TableName(contact.UserAId)).InsertOne(ctx, userContact)
		if err != nil {
			return Error(ack, err)
		}
		userContact.UserId = contact.UserAId
		ctx, cancel = NewContext()
		defer cancel()
		_, err = db.Mongo.Collection(userContact.TableName(contact.UserBId)).InsertOne(ctx, userContact)
		if err != nil {
			return Error(ack, err)
		}
	} else { //添加验证
		//查找用户
		var user = models.User{
			Username: req.Username,
		}
		ok, err := db.Mysql.Get(&user)
		if err != nil {
			return Fail(ack, "Unexpected error")
		}
		if !ok {
			return Fail(ack, "no such user")
		}

		//添加用户认证
		var contact = &models.Contact{
			UserAId:    req.Header.Uid,
			UserBId:    user.Id,
			RemarksA:   req.SetRemarks,
			CreateTime: now,
			UpdateTime: now,
			Message:    req.Msg,
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err = db.Mongo.Collection(contact.TableName()).InsertOne(ctx, contact)
		if err != nil {
			return Error(ack, err)
		}
	}

	return Success(ack)
}

func (s *ContactLogic) DeleteContact(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.DeleteContactReq)
	ack := &protocol.DeleteContactAck{}

	var userContact = &models.UserContact{}
	_id, _ := primitive.ObjectIDFromHex(req.ContactId)
	ctx, cancel := NewContext()
	defer cancel()
	var uid = userContact.UserId
	_, err := db.Mongo.Collection(userContact.TableName(req.Header.Uid)).DeleteOne(ctx, bson.M{"_id": _id})
	if err != nil {
		return Error(ack, err)
	}

	ctx, cancel = NewContext()
	defer cancel()
	_, err = db.Mongo.Collection(userContact.TableName(uid)).DeleteOne(ctx, bson.M{"user_id": req.Header.Uid})
	if err != nil {
		return Error(ack, err)
	}

	//删除联系人关系
	var contact = &models.Contact{}
	ctx, cancel = NewContext()
	defer cancel()
	_, err = db.Mongo.Collection(contact.TableName()).
		DeleteOne(ctx, bson.M{"user_id_a": req.Header.Uid, "user_id_b": uid})
	if err != nil {
		return Error(ack, err)
	}
	return Success(ack)
}

func (s *ContactLogic) UpdateContact(r Reqer) (Acker, error) {
	req, _ := r.(*protocol.UpdateContactReq)
	ack := &protocol.UpdateContactAck{}
	var userContact = &models.UserContact{}

	var ctx, cancel = NewContext()
	defer cancel()
	_id, _ := primitive.ObjectIDFromHex(req.ContactId)
	_, err := db.Mongo.Collection(userContact.TableName(req.Header.Uid)).UpdateOne(ctx, bson.M{
		"_id": _id,
	}, bson.M{"$set": bson.M{
		"remarks": req.SetRemarks,
	}})
	if err != nil {
		return Error(ack, err)
	}
	return Success(ack)
}

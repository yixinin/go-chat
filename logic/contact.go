package logic

import (
	"chat/models"
	"chat/protocol"
	"fmt"
	"go-lib/db"

	log "github.com/sirupsen/logrus"

	"github.com/go-xorm/xorm"
)

type ContactLogic struct {
}

func NewContactLogic() *ContactLogic {
	return &ContactLogic{}
}

func (s *ContactLogic) SearchUser(r Reqer) (Acker, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("SearchUser recovered", err)
		}
	}()
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
			UserId:   user.Id,
			Nickname: user.Nickname,
			Avatar:   user.Avatart,
		})
	}
	ack.Data = datas
	return Success(ack)
}

func (s *ContactLogic) AddContact(r Reqer) (Acker, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("AddContact recovered", err)
		}
	}()
	req, _ := r.(*protocol.AddContactReq)
	ack := &protocol.AddContactAck{}

	if req.AuthId != 0 { //通过验证
		ok, err := models.ApproveContact(req.AuthId, req.Add, req.SetRemarks)
		if err != nil {
			//TODO unique error
			if err == xorm.ErrNotExist { //
				return Fail(ack, "this user is already in your list")
			}
			return Error(ack, err)
		}
		if !ok {
			return Fail(ack, "add contact failed")
		}
	} else { //添加验证
		_, err := models.AddContact(req.Header.Uid, req.UserId, req.Msg, req.SetRemarks)
		if err != nil {
			return Error(ack, err)
		}
	}

	return Success(ack)
}

func (s *ContactLogic) DeleteContact(r Reqer) (Acker, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("DeleteContact recovered", err)
		}
	}()
	req, _ := r.(*protocol.DeleteContactReq)
	ack := &protocol.DeleteContactAck{}

	ok, err := models.DeleteContact(req.Header.Uid, req.ContactId)
	if err != nil {
		return Error(ack, err)
	}
	if !ok {
		return Fail(ack, "no such contact")
	}
	return Success(ack)
}

func (s *ContactLogic) UpdateContact(r Reqer) (Acker, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("UpdateContact recovered", err)
		}
	}()
	req, _ := r.(*protocol.UpdateContactReq)
	ack := &protocol.UpdateContactAck{}
	ok, err := models.DeleteContact(req.Header.Uid, req.ContactId)
	if err != nil {
		return Error(ack, err)
	}
	if !ok {
		return Fail(ack, "no such contact")
	}
	return Success(ack)
}

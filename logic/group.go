package logic

import (
	"chat/models"
	"chat/protocol"
	"go-lib/db"
)

type GroupLogic struct {
}

func (s GroupLogic) JoinGroup(r Reqer, a Acker) error {
	req, _ := r.(*protocol.JoinGroupReq)
	ack, _ := a.(*protocol.JoinGroupAck)

	var groupAuth = &models.GroupAuth{
		Token: req.GroupToken,
	}
	var ctx, cancel = NewContext()
	defer cancel()
	db.Mongo.Collection(groupAuth.TableName()).InsertOne(ctx, groupAuth)
	return Success(ack)
}

func (s *GroupLogic) Auth(r Reqer, a Acker) error {
	// req, _ := r.(*protocol.AuthGroupReq)
	ack, _ := a.(*protocol.AuthGroupAck)

	return Success(ack)
}

func (s *GroupLogic) LeaveGroup(r Reqer, a Acker) error {
	// req, _ := r.(*protocol.LeaveGroupReq)
	ack, _ := a.(*protocol.LeaveGroupAck)

	return Success(ack)
}

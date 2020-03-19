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

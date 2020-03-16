package logic

import (
	"context"
	"go-chat/cache"
	"go-chat/protocol"
	"go-lib/log"
)

type ChatServer struct {
}

func (s *ChatServer) LeaveRoom(ctx context.Context, req *protocol.LeaveRoomReq) (ack *protocol.LeaveRoomAck, err error) {
	err = cache.LeaveRoom(req.Uid)
	if err != nil {
		log.Errorf("leave room cahce del err:%v", err)
	}
	return
}

func (s *ChatServer) DiscardRoom(ctx context.Context, req *protocol.DiscardRoomReq) (ack *protocol.DiscardRoomAck, err error) {
	err = cache.DiscardRoom(req.RoomId)
	if err != nil {
		log.Errorf("discard room cahce del err:%v", err)
	}
	return
}

func (s *ChatServer) JoinRoom(ctx context.Context, req *protocol.JoinRoomReq) (ack *protocol.JoinRoomAck, err error) {
	err = cache.JoinRoom(req.User.Uid, req.RoomId)
	if err != nil {
		log.Errorf("discard room cahce del err:%v", err)
	}
	return
}

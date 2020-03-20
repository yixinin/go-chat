package handler

import (
	"chat/logic"
	"chat/protocol"
	"chat/server"
	"fmt"
	"go-lib/log"

	"github.com/davyxu/cellnet"
)

type AcceptFunc func(uid int64, v interface{})
type CloseFunc func(uid int64)

type Logic struct {
	account    *logic.AccountLogic
	contact    *logic.ContactLogic
	chat       *logic.ChatLogic
	hander     MessageHandler
	acceptFunc AcceptFunc
	closeFunc  CloseFunc
}

func NewLogic(srv server.Server) *Logic {
	return &Logic{
		hander:  EventMessageHandler,
		account: logic.NewAccountLogic(),
		chat:    logic.NewChatLogic(srv.Notify),
		contact: logic.NewContactLogic(),

		acceptFunc: srv.AcceptSess,
		closeFunc:  srv.CloseSess,
	}
}

func (s *Logic) handleMessage(sender Sender, message interface{}) {
	if sender == nil {
		return
	}
	switch msg := message.(type) {
	case *protocol.SignUpReq:
		s.hander(sender, msg, s.account.SignUp)

	case *cellnet.SessionClosed: // 会话连接断开
		fmt.Println("session closed: ", sender.ID())
	case *protocol.EchoReq:
		sender.Send(&protocol.EchoAck{
			Header: &protocol.AckHeader{
				Code: 200,
				Msg:  "Success",
			},
			Message: msg.Message,
		})
	default:
		log.Warn("no such msg", msg)
	}

}

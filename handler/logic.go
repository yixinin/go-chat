package handler

import (
	"chat/handler/iface"
	"chat/logic"
	"chat/protocol"
	"chat/server"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/davyxu/cellnet"
)

type AcceptFunc func(v *iface.Session)
type CloseFunc func(uid int64)
type AuthFunc func(*protocol.ReqHeader) bool

type Logic struct {
	account    *logic.AccountLogic
	contact    *logic.ContactLogic
	chat       *logic.ChatLogic
	hander     MessageHandler
	acceptFunc AcceptFunc
	closeFunc  CloseFunc
	authFunc   AuthFunc
}

func NewLogic(srv server.Server) *Logic {
	var s = &Logic{
		// hander:  EventMessageHandler,
		account: logic.NewAccountLogic(),
		chat:    logic.NewChatLogic(srv.Notify),
		contact: logic.NewContactLogic(),

		acceptFunc: srv.AcceptSess,
		closeFunc:  srv.CloseSess,
		authFunc:   srv.Auth,
	}
	s.hander = s.EventMessageHandler
	return s
}

func (s *Logic) handleMessage(sender iface.Sender, message interface{}) {
	if sender == nil {
		log.Error("sender is nil")
		return
	}
	switch msg := message.(type) {
	case *protocol.SignUpReq:
		s.hander(sender, msg, s.account.SignUp)
	case *protocol.SignInReq:
		s.hander(sender, msg, s.account.SignIn)
	case *protocol.SignOutReq:
		s.hander(sender, msg, s.account.SignOut)
		s.closeFunc(msg.Header.Uid)

	case *protocol.SendMessageReq:
		s.hander(sender, msg, s.chat.SendMessage)

	case *protocol.RealTimeReq:
		s.hander(sender, msg, s.chat.RealTime)
	case *protocol.GetMessageUserReq:
		s.hander(sender, msg, s.chat.GetMessageUser)
	case *protocol.GetMessageReq:
		s.hander(sender, msg, s.chat.GetUserMessage)

	case *protocol.AddContactReq:
		s.hander(sender, msg, s.contact.AddContact)
	case *protocol.GetContactListReq:
		s.hander(sender, msg, s.contact.GetContacts)

	case *cellnet.SessionAccepted:
		log.Debugln("server accepted")
	case *cellnet.SessionClosed: // 会话连接断开
		s.closeFunc(sender.ID())
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
		log.Warn("msg not handle", msg)
	}

}

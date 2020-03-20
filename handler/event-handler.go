package handler

import (
	"chat/cache"
	"chat/logic"
	"chat/protocol"
	"go-lib/log"

	"github.com/davyxu/cellnet"
)

type Eventhandler struct {
}

func (h *Eventhandler) HandleEvent(ev cellnet.Event) {
	var message = ev.Message()
	reqer, ok := message.(logic.Reqer)
	if !ok {
		return
	}
	header := reqer.GetHeader()
	if header == nil {
		return
	}
	var ack logic.Acker = nil
	var err error
	ok = h.Auth(header)
	sess := ev.Session()
	switch msg := message.(type) {
	case *protocol.SignUpReq:

	default:
		log.Warn("no such msg", msg)
	}

	if err != nil {
		log.Error(err)
	}
	if ack != nil {
		sess.Send(ack)
	}
}

func (h *Eventhandler) Auth(header *protocol.ReqHeader) bool {
	//通过header 获取token
	if header.Token == "" {
		return false
	}
	uid, err := cache.GetToken(header.Token)
	if err != nil {
		return false
	}
	header.Uid = uid
	return uid > 0
}

func AccessDeined(sess cellnet.Session, ack logic.Acker) {
	header := ack.GetHeader()
	header.Code = 401
	header.Msg = "access deined"
	sess.Send(ack)
}

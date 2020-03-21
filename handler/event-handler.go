package handler

import (
	"chat/cache"
	"chat/logic"
	"chat/protocol"
	"reflect"

	"github.com/davyxu/cellnet"
)

type Event struct {
	logic *Logic
}

func NewEvent(l *Logic) *Event {
	return &Event{
		logic: l,
	}
}

func (h *Event) HandleCallback(ev cellnet.Event) {
	var msg = ev.Message()
	reqer, ok := msg.(logic.Reqer)
	if !ok {
		return
	}
	header := reqer.GetHeader()
	if header == nil {
		header = &protocol.ReqHeader{}
		var v = reflect.Indirect(reflect.ValueOf(msg))
		v.FieldByName("Header").Set(reflect.ValueOf(header))
	} else {
		if ok := h.logic.authFunc(header); !ok {
			h.Auth(header)
		}
	}

	sess := ev.Session()
	if msg != nil {
		h.logic.handleMessage(sess, msg)
	}

}

func (h *Event) Auth(header *protocol.ReqHeader) bool {
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
	if ack == nil {
		return
	}
	header := ack.GetHeader()
	header.Code = 401
	header.Msg = "access deined"
	sess.Send(ack)
}

func (h *Event) String() string {
	return "event"
}

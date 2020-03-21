package handler

import (
	"chat/handler/iface"
	"chat/logic"
	"chat/protocol"

	log "github.com/sirupsen/logrus"
)

type Handler func(req logic.Reqer) (logic.Acker, error)

// type MessageHandler func(req logic.Reqer, ack logic.Acker) error

// type HttpHandler func(c *gin.Context, req logic.Reqer, handler Handler)

type MessageHandler func(sender iface.Sender, req logic.Reqer, handler Handler)

func (s *Logic) EventMessageHandler(sender iface.Sender, req logic.Reqer, handler Handler) {
	ack, err := handler(req)
	if err != nil {
		log.Errorf("event err:%v", err)
		log.Errorf("event req: %+v", req)
		log.Errorf("event ack: %+v", ack)
	}
	if !FillAckHeader(ack, err) {
		log.Error("ack is nil")
		return
	}

	if a, ok := ack.(*protocol.SignInAck); ok {
		if a.Header.Code == 200 && a.Token != "" { //登录成功
			s.acceptFunc(iface.NewSessoin(sender, a.Header.Uid, a.Token))
			a.Header.Uid = 0
		}
	} else {
		if a, ok := ack.(*protocol.SignUpAck); ok {
			if a.Header.Code == 200 && a.Token != "" { //登录成功
				s.acceptFunc(iface.NewSessoin(sender, a.Header.Uid, a.Token))
				a.Header.Uid = 0
			}
		}
	}

	sender.Send(ack)
}

// func HttpWMessageHandler(w http.ResponseWriter, req logic.Reqer, handler Handler) {
// 	ack, err := handler(req)
// 	if err != nil {
// 		log.Errorf("event err:%v", err)
// 		log.Errorf("event req: %+v", req)
// 		log.Errorf("event ack: %+v", ack)
// 	}
// 	if !FillHeader(ack, err) {
// 		log.Error("ack is nil")
// 		return
// 	}
// 	buf, _, err := codec.EncodeMessage(ack, nil)
// 	n, err := w.Write(buf)
// 	if err != nil {
// 		log.Error(err)
// 	}
// 	if n != len(buf) {
// 		log.Warnf("http writer not sent all, sent:%d, expect:%d", n, len(buf))
// 	}
// }

func FillAckHeader(ack logic.Acker, err error) bool {

	if ack == nil {
		return false
	}
	if err != nil {
		var header = ack.GetHeader()
		if header == nil {
			header = &protocol.AckHeader{}
		}
		if header.Code == 0 {
			header.Code = 400
		}
		if header.Msg == "" {
			header.Msg = "Unexpect error"
		}
	} else {
		var header = ack.GetHeader()
		if header == nil {
			header = &protocol.AckHeader{}
		}
		if header.Code == 0 {
			header.Code = 200
		}
		if header.Msg == "" {
			header.Msg = "Success"
		}
	}
	return true
}

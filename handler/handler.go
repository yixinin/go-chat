package handler

import (
	"chat/logic"
	"chat/protocol"
	"go-lib/log"
)

type Handler func(req logic.Reqer) (logic.Acker, error)

// type MessageHandler func(req logic.Reqer, ack logic.Acker) error

// type HttpHandler func(c *gin.Context, req logic.Reqer, handler Handler)

type MessageHandler func(sender Sender, req logic.Reqer, handler Handler)

// type HttpWHandler func(w http.ResponseWriter, req logic.Reqer, handler Handler)

// func HttpMessageHandler(c *gin.Context, req logic.Reqer, handler Handler) {
// 	if req == nil || handler == nil {
// 		return
// 	}

// 	reqHeader := req.GetHeader()
// 	if c != nil {
// 		err := c.ShouldBind(req)
// 		if err != nil {
// 			c.Status(http.StatusBadRequest)
// 			return
// 		}

// 		uid, _ := c.Get("uid")
// 		reqHeader.Uid, _ = uid.(int64)
// 		reqHeader.Token, _ = c.Cookie("token")
// 	}

// 	ack, err := handler(req)
// 	if err != nil {
// 		log.Errorf("http err:%v", err)
// 		log.Errorf("http req: %+v", req)
// 		log.Errorf("http ack: %+v", ack)

// 	}
// 	if !FillAckHeader(ack, err) {
// 		log.Error("ack is nil")
// 		return
// 	}
// 	c.JSON(http.StatusOK, ack)
// }

func (s *Logic) EventMessageHandler(sender Sender, req logic.Reqer, handler Handler) {
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
			s.acceptFunc(a.Header.Uid, sender)
			a.Header.Uid = 0
		}
	} else {
		if a, ok := ack.(*protocol.SignUpAck); ok {
			if a.Header.Code == 200 && a.Token != "" { //登录成功
				s.acceptFunc(a.Header.Uid, sender)
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

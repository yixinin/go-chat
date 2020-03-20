package handler

import (
	"chat/logic"
	"chat/protocol"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type Handler interface {
// 	Handle(c *gin.Context, relativePath string, req interface{})
// }

type MessageHandler func(req logic.Reqer, ack logic.Acker) error

type HttpHandler func(c *gin.Context, req logic.Reqer, ack logic.Acker, handler MessageHandler)

func HttpMessageHandler(c *gin.Context, req logic.Reqer, ack logic.Acker, handler MessageHandler) {
	if req == nil || handler == nil {
		return
	}

	reqHeader := req.GetHeader()
	if c != nil {
		err := c.ShouldBind(req)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		uid, _ := c.Get("uid")
		reqHeader.Uid, _ = uid.(int64)
		reqHeader.Token, _ = c.Cookie("token")
	}

	err := handler(req, ack)
	if err != nil {
		var header = ack.GetHeader()
		if header == nil {
			header = &protocol.AckHeader{}
		}
		if header.Code == 0 {
			header.Code = 400
		}
		if header.Msg == "" {
			header.Msg = "unexpect error"
		}
		c.JSON(http.StatusOK, ack)
		return
	}
	c.JSON(http.StatusOK, ack)
}

type Handler func(req logic.Reqer) (logic.Acker, error)

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

func SetHeader(c *gin.Context, header *protocol.ReqHeader) bool {
	if header == nil {
		header = &protocol.ReqHeader{}
	}
	token, err := c.Cookie("token")
	if err == nil {
		header.Token = token
	}
	if uid, ok := c.Get("uid"); ok {
		if id, ok := uid.(string); ok {
			header.Uid = id
			return true
		}
	}
	return false
}

func NotLogin(c *gin.Context) {
	c.String(http.StatusForbidden, "access deined")
}

type HttpHandler func(c *gin.Context, req logic.Reqer, ack logic.Acker, handler MessageHandler)

func HttpMessageHandler(c *gin.Context, req logic.Reqer, ack logic.Acker, handler MessageHandler) {
	err := c.ShouldBind(req)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	reqHeader := req.GetHeader()
	uid, _ := c.Get("uid")
	reqHeader.Uid, _ = uid.(string)
	reqHeader.Token, _ = c.Cookie("token")

	err = handler(req, ack)
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

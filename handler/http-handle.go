package handler

import (
	"chat/protocol"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

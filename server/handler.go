package server

import (
	"github.com/davyxu/cellnet"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	String() string
}

type HttpHandler interface {
	Handler
	HandleAll(*gin.Engine) error
}

type EventHandler interface {
	Handler
	HandleCallback(cellnet.Event)
}

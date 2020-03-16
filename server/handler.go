package server

import "github.com/gin-gonic/gin"

type Handler interface {
	String() string
}

type HttpHandler interface {
	Handler
	HandleAll(*gin.Engine) error
}

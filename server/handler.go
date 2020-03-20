package server

import (
	"net/http"

	"github.com/davyxu/cellnet"
)

type Handler interface {
	String() string
}

type HttpHandler interface {
	Handler
	Handle(http.ResponseWriter, *http.Request)
}

type EventHandler interface {
	Handler
	HandleCallback(cellnet.Event)
}

package server

import (
	"chat/handler/iface"
	"chat/protocol"
	"go-lib/hook"
)

type Server interface {
	hook.ShutdownHooker
	Init(handlers Handler) error
	Start() error
	Notify(uid int64, msg interface{}) (ok bool, err error)
	AcceptSess(v *iface.Session)
	CloseSess(uid int64)
	Auth(*protocol.ReqHeader) (ok bool)
}

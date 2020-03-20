package server

import (
	"go-lib/hook"
)

type Server interface {
	hook.ShutdownHooker
	Init(handlers Handler) error
	Start() error
	Notify(uid int64, msg interface{}) (ok bool, err error)
	AcceptSess(uid int64, v interface{})
	CloseSess(uid int64)
}

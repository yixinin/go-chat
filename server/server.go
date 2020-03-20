package server

import "go-lib/hook"

type Server interface {
	hook.ShutdownHooker
	Init(handlers ...Handler) error
	Start() error
}

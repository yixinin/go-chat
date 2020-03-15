package server

type Server interface {
	Init() error
	Start() error
	Stop() error
}

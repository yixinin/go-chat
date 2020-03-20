package grpc

import (
	"chat/logic"
	"chat/protocol"
	"chat/server"
)

type Config struct {
	Port string
}

type GrpcServer struct {
	server protocol.RoomServiceServer
	config *Config
}

func NewGrpcServer(c *Config) server.Server {
	var s = &GrpcServer{
		config: c,
	}

	return s
}

func (s *GrpcServer) Init(handlers server.Handler) error {

	return nil
}

func (s *GrpcServer) Start() error {
	return nil
}
func (s *GrpcServer) Stop() error {
	return nil
}

func (s *GrpcServer) Shutdown() {

}

func (s *GrpcServer) GetNotifyFunc() logic.NotifyFunc {
	return nil
}

func (s *GrpcServer) AcceptSess(uid int64, sess interface{}) {

}

func (s *GrpcServer) CloseSess(uid int64) {

}

func (s *GrpcServer) Notify(uid int64, msg interface{}) (ok bool, err error) {
	return false, nil
}

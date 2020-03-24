package grpc

import (
	"chat/handler/iface"
	"chat/logic"
	"chat/protocol"
	"chat/server"
)

type Config struct {
	Addr string
	Host string
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

func (s *GrpcServer) AcceptSess(sess *iface.Session) {

}

func (s *GrpcServer) CloseSess(uid int64) {

}

func (s *GrpcServer) Notify(uid int64, msg interface{}) (ok bool, err error) {
	return false, nil
}

func (s *GrpcServer) Auth(header *protocol.ReqHeader) bool {
	return false
}

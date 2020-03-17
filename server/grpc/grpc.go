package grpc

import (
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

func (s *GrpcServer) Init(handlers ...server.Handler) error {

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

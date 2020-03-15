package main

import (
	"go-chat/server"
	"go-lib/registry"
	"time"

	"go-chat/config"
)

type Service struct {
	WsServer   server.Server
	TcpServer  server.Server
	HttpServer server.Server
	GrpcServer server.Server

	Address string
	ID      string

	Registry    registry.Registry
	serviceInfo *registry.Service
	config      *config.Config
}

func (s *Service) Init() {
	s.Registry.Init(registry.Addrs(s.config.EtcdAddr...))
	s.TcpServer.Init()
	s.HttpServer.Init()
	s.WsServer.Init()
	s.GrpcServer.Init()

	s.serviceInfo = &registry.Service{
		Name:    "chat",
		Version: "1.0",
	}
}

func (s *Service) Start() {

	s.TcpServer.Start()
	s.HttpServer.Start()
	s.WsServer.Start()
	s.GrpcServer.Start()

	//注册服务
	s.Registry.Register(&registry.Service{
		Name:    "live-chat.chat",
		Version: "v2",
		Nodes: []*registry.Node{
			{Id: s.ID, Address: s.Address},
		},
	}, registry.RegisterTTL(time.Minute))
}

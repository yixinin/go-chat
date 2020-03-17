package main

import (
	"fmt"
	"go-chat/handler"
	"go-chat/logic"
	"go-chat/protocol"
	"go-chat/server"
	"go-chat/server/grpc"
	"go-chat/server/http"
	"go-lib/ip"
	"go-lib/log"
	"go-lib/registry"
	"go-lib/registry/etcd"
	"net"
	"time"

	"go-lib/utils"

	ggrpc "google.golang.org/grpc"

	"go-chat/config"
)

type Service struct {
	WsServer   server.Server
	TcpServer  server.Server
	HttpServer server.Server
	GrpcServer server.Server
	gs         *ggrpc.Server

	watcher registry.Watcher

	Registry registry.Registry
	config   *config.Config
}

func NewService(c *config.Config) *Service {

	var regist = etcd.NewRegistry()
	watcher, err := regist.Watch()
	if err != nil {
		log.Errorf("watch err:%v", err)
	}

	var s = &Service{
		HttpServer: http.NewHttpServer(c.HttpConfig),
		GrpcServer: grpc.NewGrpcServer(c.GrpcConfig),
		config:     c,
		Registry:   regist,
		watcher:    watcher,
		gs:         ggrpc.NewServer(),
	}
	protocol.RegisterChatServiceServer(s.gs, &logic.ChatServer{})
	return s
}

func (s *Service) Init() {

	if s.HttpServer != nil {
		var account = handler.NewAccountHandler()
		var message = handler.NewMessageHandler(s.watcher)
		s.HttpServer.Init(account, message)
	}
	if s.TcpServer != nil {
		s.TcpServer.Init()
	}
	if s.WsServer != nil {
		s.WsServer.Init()
	}
	if s.GrpcServer != nil {
		s.GrpcServer.Init()
	}

	s.Registry.Init(
		registry.Addrs(s.config.EtcdAddr...),
	)

}

func (s *Service) Start() error {

	if s.TcpServer != nil {
		go s.TcpServer.Start()
	}

	if s.WsServer != nil {
		go s.WsServer.Start()
	}
	if s.HttpServer != nil {
		go s.HttpServer.Start()
	}

	if s.GrpcServer != nil {
		go s.GrpcServer.Start()
	}

	var listen, err = net.Listen("tcp", fmt.Sprintf(":%s", s.config.GrpcConfig.Port))
	if err != nil {
		return err
	}
	s.gs.Serve(listen)
	//注册服务
	s.Registry.Register(&registry.Service{
		Name:    "live-chat.chat",
		Version: "v1.0",
		Nodes: []*registry.Node{
			{Id: utils.UUID(), Address: ip.LocalIP},
		},
	}, registry.RegisterTTL(time.Minute))
	return nil
}

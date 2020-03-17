package main

import (
	"chat/handler"
	"chat/logic"
	"chat/protocol"
	"chat/server"
	"chat/server/grpc"
	"chat/server/http"
	"fmt"
	"go-lib/db"
	"go-lib/ip"
	"go-lib/log"
	"go-lib/registry"
	"go-lib/registry/etcd"
	"go-lib/utils"
	"net"
	"time"

	ggrpc "google.golang.org/grpc"

	"chat/config"
)

type Service struct {
	WsServer   server.Server
	TcpServer  server.Server
	HttpServer server.Server
	GrpcServer server.Server
	gs         *ggrpc.Server

	watcher registry.Watcher

	Registry        registry.Registry
	RegistrtService *registry.Service
	config          *config.Config

	stop chan bool
}

func NewService(c *config.Config) *Service {

	var s = &Service{
		HttpServer: http.NewHttpServer(c.HttpConfig),
		GrpcServer: grpc.NewGrpcServer(c.GrpcConfig),
		config:     c,
		Registry:   etcd.NewRegistry(),
		gs:         ggrpc.NewServer(),
		stop:       make(chan bool),
	}
	protocol.RegisterChatServiceServer(s.gs, &logic.ChatServer{})
	return s
}

func (s *Service) Init() {

	db.InitMongo(s.config.Mongo)
	db.InitRedis(s.config.Redis)

	s.Registry.Init(
		registry.Addrs(s.config.EtcdAddr...),
	)

	if s.HttpServer != nil {
		var account = handler.NewAccountHandler()
		var message = handler.NewMessageHandler(s.Registry)
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
	go func() {
		err = s.gs.Serve(listen)
		if err != nil {
			log.Error(err)
		}
	}()

	//注册服务
	var srv = &registry.Service{
		Name:    "live-chat.chat",
		Version: "v1.0",
		Nodes: []*registry.Node{
			{Id: utils.UUID(), Address: ip.GrpcAddr(s.config.GrpcConfig.Port)},
		},
	}
	s.RegistrtService = srv
	err = s.Registry.Register(srv, registry.RegisterTTL(5*time.Second))
	if err != nil {
		log.Error(err)
	}

	go s.KeepAlive()
	return nil
}

func (s *Service) KeepAlive() {
	for {
		select {
		case <-s.stop:
			err := s.Registry.Deregister(s.RegistrtService)
			if err != nil {
				log.Error(err)
			}
			return
		default:
		}
	}

}

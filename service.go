package main

import (
	"chat/handler"
	"chat/logic"
	"chat/protocol"
	"chat/server"
	"chat/server/grpc"
	"chat/server/http"
	"chat/server/tcp"
	"chat/server/ws"
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
		TcpServer:  tcp.NewTcpServer(c.TcpConfig),
		WsServer:   ws.NewWsServer(c.WsConfig),
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
	db.InitMysql(s.config.Mysql)

	s.Registry.Init(
		registry.Addrs(s.config.EtcdAddr...),
	)

	if s.HttpServer != nil {
		var logic = handler.NewLogic(s.HttpServer)
		s.HttpServer.Init(handler.NewHttp(logic))
	}
	if s.TcpServer != nil {
		var logic = handler.NewLogic(s.TcpServer)
		s.TcpServer.Init(handler.NewEvent(logic))
	}
	if s.WsServer != nil {
		var logic = handler.NewLogic(s.WsServer)
		s.WsServer.Init(handler.NewEvent(logic))
	}
	if s.GrpcServer != nil {
		// var logic = handler.NewLogic()
		s.GrpcServer.Init(handler.NewGrpcHandler())
	}

}

func (s *Service) Start() error {

	if s.TcpServer != nil {
		err := s.TcpServer.Start()
		if err != nil {
			log.Error(err)
		}
	}

	if s.WsServer != nil {
		err := s.WsServer.Start()
		if err != nil {
			log.Error(err)
		}
	}
	if s.HttpServer != nil {
		go s.HttpServer.Start()
	}

	if s.GrpcServer != nil {
		go s.GrpcServer.Start()
	}

	var listen, err = net.Listen("tcp", s.config.GrpcConfig.Addr)
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
			{Id: utils.UUID(), Address: ip.GetAddr(s.config.GrpcConfig.Addr)},
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

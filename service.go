package main

import (
	"chat/handler"
	"chat/logic"
	"chat/models"
	"chat/protocol"
	"chat/server"
	"chat/server/grpc"
	"chat/server/http"
	"chat/server/tcp"
	"chat/server/ws"
	"go-lib/db"
	"go-lib/pool"
	"go-lib/registry"
	"go-lib/registry/etcd"
	"go-lib/utils"
	"net"
	"time"

	log "github.com/sirupsen/logrus"

	ggrpc "google.golang.org/grpc"

	"chat/config"
)

const (
	VoipServiceName = "live-chat.voip"
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

	models.SyncTables()

	s.Registry.Init(
		registry.Addrs(s.config.EtcdAddr...),
	)

	var notifys = make([]logic.NotifyFunc, 0, 2)
	if s.TcpServer != nil {
		notifys = append(notifys, s.TcpServer.Notify)
	}
	if s.WsServer != nil {
		notifys = append(notifys, s.WsServer.Notify)
	}

	if s.HttpServer != nil {
		var logic = handler.NewLogic(s.HttpServer, notifys...)
		s.HttpServer.Init(handler.NewHttp(logic))
	}
	if s.TcpServer != nil {
		var logic = handler.NewLogic(s.TcpServer, notifys...)
		s.TcpServer.Init(handler.NewEvent(logic))
	}
	if s.WsServer != nil {
		var logic = handler.NewLogic(s.WsServer, notifys...)
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
			{Id: utils.UUID(), Address: s.config.GrpcConfig.Host + s.config.GrpcConfig.Addr},
		},
	}
	s.RegistrtService = srv
	err = s.Registry.Register(srv, registry.RegisterTTL(5*time.Second))
	if err != nil {
		log.Error(err)
	}

	watcher, err := s.Registry.Watch(registry.WatchService(VoipServiceName))
	if err != nil {
		log.Error(err)
		return err
	}

	go s.Watch(watcher)
	return nil
}

func (s *Service) Watch(watcher registry.Watcher) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("etcd watcher recoverd", err)
		}
	}()

	//添加初始node
	services, err := s.Registry.GetService(VoipServiceName)
	if err != nil {
		log.Error(err)
	} else {
		if services != nil {
			for _, srv := range services {
				if srv == nil || srv.Nodes == nil {
					continue
				}
				for _, node := range srv.Nodes {
					if node == nil {
						continue
					}
					var addr = node.Address
					if addr[0] == ':' {
						addr = "voip" + addr
					}
					pool.DefaultGrpcConnPool.AddNode(addr)
					log.Infof("add node %s :%s", srv.Name, addr)
				}
			}
		}

	}
FOR:
	for {
		select {
		case <-s.stop:
			log.Info("exit ...")
			err := s.Registry.Deregister(s.RegistrtService)
			if err != nil {
				log.Error(err)
			}
			return
		default:
			res, err := watcher.Next()
			if err != nil || res == nil || res.Service == nil || res.Service.Nodes == nil {
				if err != nil {
					log.Error(err)
				}
				continue FOR
			}
			if res.Service.Name == s.RegistrtService.Name {
				continue
			}
			var name = res.Service.Name
			switch res.Action {
			case registry.Create.String():
				for _, node := range res.Service.Nodes {
					if node == nil {
						continue FOR
					}
					var addr = node.Address
					if addr[0] == ':' {
						addr = "voip" + addr
					}
					pool.DefaultGrpcConnPool.AddNode(addr)
					log.Infof("----new node %s :%s", name, addr)
				}
			case registry.Update.String():
				for _, node := range res.Service.Nodes {
					if node == nil {
						continue FOR
					}
					var addr = node.Address
					if addr[0] == ':' {
						addr = "voip" + addr
					}
					pool.DefaultGrpcConnPool.AddNode(addr)
					log.Infof("----update node %s :%s", name, addr)
				}
			case registry.Delete.String():
				for _, node := range res.Service.Nodes {
					if node == nil {
						continue FOR
					}
					var addr = node.Address
					if addr[0] == ':' {
						addr = "voip" + addr
					}
					pool.DefaultGrpcConnPool.AddNode(addr)
					log.Infof("----del node %s :%s", name, addr)
				}
			default:
				for _, node := range res.Service.Nodes {
					if node == nil {
						continue FOR
					}
					log.Infof("----not cased, %s node %s :%s", res.Action, name, node.Address)
				}
			}

		}
	}

}

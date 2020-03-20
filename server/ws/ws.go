package ws

import (
	"chat/server"
	"errors"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/gorillaws"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/gorillaws"
)

type Config struct {
	Addr string
}

type WsServer struct {
	config  *Config
	handler server.EventHandler

	peer  cellnet.GenericPeer
	queue cellnet.EventQueue
}

func NewWsServer(c *Config) *WsServer {
	return &WsServer{
		config: c,
	}
}
func (s *WsServer) Init(handlers ...server.Handler) error {
	if len(handlers) > 0 {
		handler, ok := handlers[0].(server.EventHandler)
		if ok {
			s.handler = handler
		}
	}
	return nil
}
func (s *WsServer) Start() error {
	if s.handler == nil {
		return errors.New("no ws event handlers ")
	}
	s.queue = cellnet.NewEventQueue()

	// 创建一个服务器的接受器(Acceptor)，接受客户端的连接
	s.peer = peer.NewGenericPeer("gorillaws.Acceptor", "server", s.config.Addr, s.queue)

	// 将接受器Peer与gorillaws.ltv的处理器绑定，并设置事件处理回调
	// gorillaws.ltv处理器负责处理消息收发，使用私有的封包格式以及日志，RPC等处理
	proc.BindProcessorHandler(s.peer, "gorillaws.ltv", s.handler.HandleCallback)

	// 启动Peer，服务器开始侦听
	s.peer.Start()

	// 开启事件队列，开始处理事件，此函数不阻塞
	s.queue.StartLoop()
	return nil
}

func (s *WsServer) Shutdown() {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	s.peer.Stop()
	s.queue.StopLoop()
}

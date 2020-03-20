package ws

import (
	"chat/logic"
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

	// notify logic.NotifyFunc

	peer  cellnet.GenericPeer
	queue cellnet.EventQueue

	users map[int64]cellnet.Session
}

func NewWsServer(c *Config) server.Server {
	var queue = cellnet.NewEventQueue()
	var s = &WsServer{
		config: c,
		queue:  queue,
		peer:   peer.NewGenericPeer("gorillaws.Acceptor", "server", c.Addr, queue),
		users:  make(map[int64]cellnet.Session, 100),
	}
	return s
}
func (s *WsServer) Init(handler server.Handler) error {
	h, ok := handler.(server.EventHandler)
	if ok {
		s.handler = h
	}
	return nil
}
func (s *WsServer) Start() error {
	if s.handler == nil {
		return errors.New("no ws event handlers ")
	}
	// s.queue = cellnet.NewEventQueue()

	// // 创建一个服务器的接受器(Acceptor)，接受客户端的连接
	// s.peer = peer.NewGenericPeer("gorillaws.Acceptor", "server", s.config.Addr, s.queue)

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

func (s *WsServer) GetNotifyFunc() logic.NotifyFunc {
	return s.Notify
}

func (s *WsServer) Notify(uid int64, msg interface{}) (ok bool, err error) {
	if sess, ok := s.users[uid]; ok {
		sess.Send(msg)
		return true, nil
	}
	return false, nil
}

func (s *WsServer) AcceptSess(uid int64, v interface{}) {
	if old, ok := s.users[uid]; ok {
		old.Close()
		delete(s.users, uid)
	}
	if sess, ok := v.(cellnet.Session); ok {
		s.users[uid] = sess
	}
}

func (s *WsServer) CloseSess(uid int64) {
	if _, ok := s.users[uid]; ok {
		delete(s.users, uid)
	} else {
		//通过ID删除 被动断线
		for k, v := range s.users {
			if v.ID() == uid {
				delete(s.users, k)
				return
			}
		}
	}
}

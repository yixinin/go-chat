package tcp

import (
	"chat/logic"
	"chat/server"
	"errors"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/tcp"
)

type Config struct {
	Addr string
}

type TcpServer struct {
	config  *Config
	handler server.EventHandler

	peer  cellnet.GenericPeer
	queue cellnet.EventQueue

	users map[int64]cellnet.Session
}

func NewTcpServer(c *Config) server.Server {
	var queue = cellnet.NewEventQueue()
	var s = &TcpServer{
		config: c,
		queue:  queue,
		peer:   peer.NewGenericPeer("tcp.Acceptor", "server", c.Addr, queue),
		users:  make(map[int64]cellnet.Session, 100),
	}
	return s
}
func (s *TcpServer) Init(handler server.Handler) error {
	h, ok := handler.(server.EventHandler)
	if ok {
		s.handler = h
	}
	return nil
}
func (s *TcpServer) Start() error {
	if s.handler == nil {
		return errors.New("no tcp event handlers ")
	}
	// s.queue = cellnet.NewEventQueue()

	// // 创建一个服务器的接受器(Acceptor)，接受客户端的连接
	// s.peer = peer.NewGenericPeer("tcp.Acceptor", "server", s.config.Addr, s.queue)

	// 将接受器Peer与tcp.ltv的处理器绑定，并设置事件处理回调
	// tcp.ltv处理器负责处理消息收发，使用私有的封包格式以及日志，RPC等处理
	proc.BindProcessorHandler(s.peer, "tcp.ltv", s.handler.HandleCallback)

	// 启动Peer，服务器开始侦听
	s.peer.Start()

	// 开启事件队列，开始处理事件，此函数不阻塞
	s.queue.StartLoop()
	return nil
}

func (s *TcpServer) Shutdown() {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	s.peer.Stop()
	s.queue.StopLoop()
}

func (s *TcpServer) GetNotifyFunc() logic.NotifyFunc {
	return s.Notify
}

func (s *TcpServer) Notify(uid int64, msg interface{}) (ok bool, err error) {
	if sess, ok := s.users[uid]; ok {
		sess.Send(msg)
		return true, nil
	}
	return false, nil
}

func (s *TcpServer) AcceptSess(uid int64, v interface{}) {
	if old, ok := s.users[uid]; ok {
		old.Close()
		delete(s.users, uid)
	}
	if sess, ok := v.(cellnet.Session); ok {
		s.users[uid] = sess
	}
}

func (s *TcpServer) CloseSess(uid int64) {
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

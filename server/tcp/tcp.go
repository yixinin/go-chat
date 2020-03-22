package tcp

import (
	"chat/handler/iface"
	"chat/logic"
	"chat/protocol"
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

	users  map[int64]*iface.Session
	tokens map[string]int64
	sess   map[int64]int64
}

func NewTcpServer(c *Config) server.Server {
	var queue = cellnet.NewEventQueue()
	var s = &TcpServer{
		config: c,
		queue:  queue,
		peer:   peer.NewGenericPeer("tcp.Acceptor", "server", c.Addr, queue),
		users:  make(map[int64]*iface.Session, 100),
		tokens: make(map[string]int64, 100),
		sess:   make(map[int64]int64, 100),
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

func (s *TcpServer) AcceptSess(sess *iface.Session) {
	if old, ok := s.users[sess.Uid]; ok {
		old.Close()
		delete(s.users, sess.Uid)
		delete(s.tokens, old.Token)
		delete(s.sess, old.ID())
	}
	s.users[sess.Uid] = sess
	s.tokens[sess.Token] = sess.Uid
	s.sess[sess.ID()] = sess.Uid
}

func (s *TcpServer) CloseSess(v int64) {
	if u, ok := s.users[v]; ok {
		delete(s.users, v)
		delete(s.sess, u.ID())
		delete(s.tokens, u.Token)
	} else {
		//通过ID删除 被动断线
		if uid, ok := s.sess[v]; ok {
			if u, ok := s.users[uid]; ok {
				delete(s.users, uid)
				delete(s.sess, u.ID())
				delete(s.tokens, u.Token)
			}
		}

	}
}

func (s *TcpServer) Auth(header *protocol.ReqHeader) bool {
	if header == nil {
		return false
	}
	_, ok := s.tokens[header.Token]
	return ok
}

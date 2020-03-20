package http

import (
	"chat/logic"
	"chat/server"
	"context"
	"flag"
	"go-lib/log"
	"net/http"
	"time"
)

var (
	DebugMode = flag.Bool("debug", true, "-debug")
)

type Config struct {
	Addr   string
	Router string
}

type httpServer struct {
	handler server.HttpHandler
	// engine  *gin.Engine
	hs     *http.Server
	config *Config
	users  map[int64]int64
}

func NewHttpServer(c *Config) server.Server {
	var addr = "0.0.0.0:8080"
	if c != nil && c.Addr != "" {
		addr = c.Addr
	}
	var s = &httpServer{
		hs: &http.Server{
			Addr: addr,
		},
		config: c,
		users:  make(map[int64]int64, 100),
	}
	return s
}

func (s *httpServer) Start() (err error) {

	if s.handler != nil {
		http.HandleFunc(s.config.Router, s.handler.Handle)
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("http ListenAndServe panic! recover:%v", err)
			}
		}()
		if err := s.hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("listen: %s\n", err)
		}
	}()
	return nil
}

func (s *httpServer) Init(handler server.Handler) error {
	// s.engine = gin.Default()
	// if *DebugMode {
	// 	gin.SetMode(gin.DebugMode)
	// } else {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	// s.engine.Use(middleware.SetUid)
	if h, ok := handler.(server.HttpHandler); ok {
		s.handler = h
	}

	return nil
}

func (s *httpServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.hs.Shutdown(ctx)
}

func (s *httpServer) Stop() error {

	return nil
}

func (s *httpServer) GetNotifyFunc() logic.NotifyFunc {
	return nil
}

func (s *httpServer) AcceptSess(uid int64, v interface{}) {
	s.users[uid] = time.Now().Unix()
}

func (s *httpServer) CloseSess(uid int64) {
	if _, ok := s.users[uid]; ok {
		delete(s.users, uid)
	}
}

func (s *httpServer) Notify(uid int64, msg interface{}) (ok bool, err error) {
	return false, nil
}

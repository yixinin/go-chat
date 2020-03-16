package http

import (
	"context"
	"flag"
	"go-chat/server"
	"go-lib/log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	DebugMode = flag.Bool("debug", true, "-debug")
)

type Config struct {
	Addr string
}

type httpServer struct {
	handlers []server.HttpHandler
	engine   *gin.Engine
	hs       *http.Server
	config   *Config
}

func NewHttpServer(c *Config) server.Server {
	var addr = "0.0.0.0:8080"
	if c != nil && c.Addr != "" {
		addr = c.Addr
	}
	var s = &httpServer{
		handlers: make([]server.HttpHandler, 0, 2),
		hs: &http.Server{
			Addr: addr,
		},
		config: c,
	}
	return s
}

func (s *httpServer) Start() (err error) {

	for _, h := range s.handlers {
		err := h.HandleAll(s.engine)
		if err != nil {
			return err
		}
	}
	s.hs.Handler = s.engine
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

func (s *httpServer) Init(handlers ...server.Handler) error {
	s.engine = gin.Default()
	if *DebugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	for _, h := range handlers {
		if handler, ok := h.(server.HttpHandler); ok {
			s.handlers = append(s.handlers, handler)
		} else {
			log.Warnf("handler %s is not http handler", h.String())
		}
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

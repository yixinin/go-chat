package main

import (
	"flag"
	"go-lib/ip"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"

	"chat/config"
)

var confPath = flag.String("conf", "config/app.yaml", "-conf=xxx")

func main() {
	var conf, err = config.GetConfig(*confPath)
	if err != nil {
		log.Panicf("config not found, err:%v", err)
	}
	showIP("grpc", conf.GrpcConfig.Host+conf.GrpcConfig.Addr)
	showIP("tcp", conf.TcpConfig.Addr)
	showIP("ws", conf.WsConfig.Addr)
	showIP("http", conf.HttpConfig.Addr)
	var service = NewService(conf)
	service.Init()
	service.Start()

	//监听退出信号
	c := make(chan os.Signal)
	//监听所有信号
	signal.Notify(c)

	for {
		select {
		case sig := <-c:
			switch sig {
			case os.Interrupt:
				close(service.stop)
				return
			}
		}
	}
}

func showIP(s string, port string) {
	log.Infof("%s listen on: %s", s, ip.GetAddr(port))
}

package main

import (
	"flag"
	"go-lib/ip"
	"go-lib/log"
	"os"
	"os/signal"

	"chat/config"
)

var confPath = flag.String("conf", "config/app.yaml", "-conf=xxx")

func main() {
	var conf, err = config.GetConfig(*confPath)
	if err != nil {
		log.Panicf("config not found, err:%v", err)
	}
	showIP(conf.GrpcConfig.Port)
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
				service.stop <- true
				return
			}
		}
	}
}

func showIP(port string) {
	log.Infof("本机IP:%s", ip.GrpcAddr(port))
}

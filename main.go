package main

import (
	"flag"
	"go-lib/log"

	"go-chat/config"
)

var confPath = flag.String("conf", "config/app.yaml", "-conf=xxx")

func main() {
	var conf, err = config.GetConfig(*confPath)
	if err != nil {
		log.Panicf("config not found, err:%v", err)
	}
	var service = NewService(conf)
	service.Init()
	service.Start()
}

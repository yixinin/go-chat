package config

import (
	"chat/server/grpc"
	"chat/server/http"
	"chat/server/tcp"
	"chat/server/ws"
	"go-lib/db"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

type Config struct {
	EtcdAddr   []string        `yaml:"etcd"`
	GrpcConfig *grpc.Config    `yaml:"grpc"`
	HttpConfig *http.Config    `yaml:"http"`
	TcpConfig  *tcp.Config     `yaml:"tcp"`
	WsConfig   *ws.Config      `yaml:"ws"`
	Mongo      *db.MongoConfig `yaml:"mongo"`
	Redis      *db.RedisConfig `yaml:"redis"`
	Mysql      *db.MysqlConfig `yaml:"mysql"`
}

func GetConfig(p string) (*Config, error) {
	var c Config
	yamlFile, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(yamlFile, &c)
	return &c, nil
}
func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// log.SetReportCaller(true)
}

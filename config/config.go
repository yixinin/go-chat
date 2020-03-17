package config

import (
	"chat/server/grpc"
	"chat/server/http"
	"go-lib/db"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	EtcdAddr   []string        `yaml:"etcd"`
	GrpcConfig *grpc.Config    `yaml:"grpc"`
	HttpConfig *http.Config    `yaml:"http"`
	Mongo      *db.MongoConfig `yaml:"mongo"`
	Redis      *db.RedisConfig `yaml:"redis"`
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

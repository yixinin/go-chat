package config

import "go-chat/server/http"

type Config struct {
	EtcdAddr []string

	HttpConfig *http.Config
}

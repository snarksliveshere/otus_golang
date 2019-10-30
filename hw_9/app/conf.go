package app

import (
	"github.com/caarlos0/env"
	"sync"
)

type config struct {
	ListenIp   string `env:"LISTEN_IP" envDefault:"127.0.0.1"`
	ListenPort string `env:"LISTEN_PORT" envDefault:"3003"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"info"`
}

var (
	confOnce sync.Once
	conf     *config
)

func Conf() *config {
	confOnce.Do(func() {
		conf = &config{}
		err := env.Parse(conf)
		if err != nil {
			panic(err)
		}
	})
	return conf
}

func ListenAddr(conf *config) string {
	return conf.ListenIp + ":" + conf.ListenPort
}

package app

import (
	"fmt"
	"github.com/caarlos0/env"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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
	cf, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	yc := &YamlFileConfig{}
	yc.parse(cf)
	fmt.Println(yc.ListenIp)

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

type YamlFileConfig struct {
	ListenIp string `yaml:"listenIp"`
}

func (y *YamlFileConfig) parse(sb []byte) {
	err := yaml.Unmarshal(sb, y)
	if err != nil {
		log.Fatal(err.Error())
	}
}

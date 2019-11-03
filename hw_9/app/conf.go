package app

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

type config struct {
	ListenIp   string `yaml:"listen_ip"`
	ListenPort string `yaml:"listen_port"`
	LogLevel   string `yaml:"log_level"`
}

var (
	confOnce sync.Once
	conf     *config
)

func Conf() *config {
	confOnce.Do(func() {
		conf = &config{}
		conf.parse()
	})
	return conf
}

func ListenAddr(conf *config) string {
	return conf.ListenIp + ":" + conf.ListenPort
}

func (conf *config) parse() {
	cf, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = yaml.Unmarshal(cf, conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}

package app

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

type Config struct {
	ListenIp   string `yaml:"listen_ip"`
	ListenPort string `yaml:"listen_port"`
	LogLevel   string `yaml:"log_level"`
}

var (
	confOnce sync.Once
	conf     *Config
)

func Conf(path string) *Config {
	confOnce.Do(func() {
		conf = &Config{}
		conf.parse(path)
	})
	return conf
}

func ListenAddr(conf *Config) string {
	return conf.ListenIp + ":" + conf.ListenPort
}

func (conf *Config) parse(path string) {
	cf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = yaml.Unmarshal(cf, conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	ListenIp       string `yaml:"listen_ip"`
	ListenPort     string `yaml:"listen_port"`
	LogLevel       string `yaml:"log_level"`
	RabbitPort     string `yaml:"rabbit_port"`
	RabbitUser     string `yaml:"rabbit_user"`
	RabbitPassword string `yaml:"rabbit_password"`
}

func CreateConfig(path string) *Config {
	cf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	conf := &Config{}
	err = yaml.Unmarshal(cf, conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conf
}

func (conf *Config) ListenAddr() string {
	return conf.ListenIp + ":" + conf.ListenPort
}

package main

import (
	"flag"
	"github.com/snarskliveshere/otus_golang/hw_8/cmd/web"
)

var (
	pathConfig string
)

const (
	confFile = "./config/config.yaml"
)

func init() {
	flag.StringVar(&pathConfig, "config", confFile, "path config")
}

func main() {
	flag.Parse()
	web.Server(pathConfig)
}

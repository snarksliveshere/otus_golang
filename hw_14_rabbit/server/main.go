package main

import (
	"flag"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/cmd/grpc"
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
	grpc.Server(pathConfig)
}

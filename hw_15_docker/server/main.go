package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/cmd/grpc"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/cmd/web"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/config"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/pkg/logger/logrus"
	"log"
)

func main() {
	var conf config.AppConfig
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")
	logg := logrus.CreateLogrusLog(conf.LogLevel)
	web.Server(logg, &conf)
	grpc.Server(logg, &conf)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

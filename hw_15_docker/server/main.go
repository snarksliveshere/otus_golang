package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/cmd/grpc"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/cmd/web"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/config"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/pkg/logger/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//time.Sleep(30 * time.Second)
	var conf config.AppConfig
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")
	logg := logrus.CreateLogrusLog(conf.LogLevel)
	logg.Info("start web server")
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	go func() { web.Server(logg, &conf) }()
	logg.Info("start grpc server")
	go func() { grpc.Server(logg, &conf) }()
	<-stopCh
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

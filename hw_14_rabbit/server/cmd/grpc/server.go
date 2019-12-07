package grpc

import (
	"fmt"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/config"
	pg_repository2 "github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/internal/interfaces/repositories/pg_repository"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/proto"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/tools/logger"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type ServerCalendar struct {
}

var (
	log     *logger.Logger
	storage *pg_repository2.Storage
)

func Server(path string) {
	conf := config.CreateConfig(path)
	log = logger.CreateLogrusLog(conf)

	storage = pg_repository2.CreateStorageInstance(log, conf)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	webApi(conf)
	<-stopCh
}

func webApi(conf *config.Config) {
	listenAddr := conf.ListenAddr()
	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen addr: %s, error: %v\n", listenAddr, err.Error())
	}

	grpcServer := grpc.NewServer()
	proto.RegisterEventServiceServer(grpcServer, ServerCalendar{})
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println(err.Error())
	}

	log.Infof("Run grpc server on: %s\n", listenAddr)
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.

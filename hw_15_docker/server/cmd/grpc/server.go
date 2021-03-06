package grpc

import (
	"fmt"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/config"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/internal/interfaces/repositories/pg_repository"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/pkg/logger/logrus"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/proto"
	"google.golang.org/grpc"
	"net"
)

type ServerCalendar struct {
}

var (
	log     *logrus.Logger
	storage *pg_repository.Storage
)

func Server(logg *logrus.Logger, conf *config.AppConfig) {
	log = logg
	storage = pg_repository.CreateStorageInstance(log, conf)
	goGRPC(conf)
}

func goGRPC(conf *config.AppConfig) {
	listenAddr := conf.ListenIP + ":" + conf.GRPCPort
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

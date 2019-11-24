package web

import (
	"context"
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/cmd/inmem"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/config"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/pkg"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/proto"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	log     *pkg.Logger
	storage *inmem.Storage
)

func Server(path string) {
	conf := config.CreateConfig(path)
	log = pkg.CreateLog(conf)

	storage = inmem.CreateStorageInstance(log)

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
	proto.RegisterCreateEventServiceServer(grpcServer, ServerCalendar{})
	err = grpcServer.Serve(listen)
	if err != nil {
		fmt.Println(err.Error())
	}

	log.Infof("Run grpc server on: %s\n", listenAddr)
}

type ServerCalendar struct {
}

func (s ServerCalendar) SendCreateEventMessage(ctx context.Context, msg *proto.CreateEventMsg) (*proto.CreateEventMessage, error) {
	fmt.Println("find2")
	title, desc, day, err := validCreateEventHandler(msg.Title, msg.Description, msg.Date)
	if err != nil {
		return nil, err
	}
	rec, dt, c, err := storage.AddRecord(title, desc, day)
	reply := proto.CreateEventMessage{}

	if err != nil {
		reply.Status = "error"
		reply.Error = err.Error()
	}
	fmt.Println(rec, dt, c)
	reply.Status = "success"
	return &reply, nil

}

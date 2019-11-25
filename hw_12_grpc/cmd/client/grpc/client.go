package main

import (
	"context"
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	cc, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func() { _ = cc.Close() }()

	switch expr := os.Args[1]; expr {
	case "create-event":
		sendCreateMessage(ctx, cc)
	default:
		fmt.Println("bad route")
	}

}

func sendCreateMessage(ctx context.Context, cc *grpc.ClientConn) *proto.CreateEventResponseMessage {
	c := proto.NewCreateEventServiceClient(cc)
	message := proto.CreateEventRequestMessage{
		Title:       "some title",
		Description: "some description",
		Date:        "2019-11-01",
	}
	msg, err := c.SendCreateEventMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}
	if msg != nil {
		fmt.Printf("error:%v status:%v\n, record: %#v, id %v", msg.Error, msg.Status, msg.Record, msg.Record.Id)
	}

	return msg
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.

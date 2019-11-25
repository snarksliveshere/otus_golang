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

	msgCreateEvent := Dummy{createEventReq: proto.CreateEventRequestMessage{
		Title:       "Some_title",
		Description: "Some_description",
		Date:        "2019-11-01",
	}}
	switch expr := os.Args[1]; expr {
	case "create-event":
		sendCreateEventMessage(ctx, cc, msgCreateEvent.createEventReq)
	case "delete-event":
		rec := sendCreateEventMessage(ctx, cc, msgCreateEvent.createEventReq)
		msgDeleteEvent := Dummy{deleteEventReq: proto.DeleteEventRequestMessage{
			EventId: rec.Record.Id,
		}}
		sendDeleteEventMessage(ctx, cc, msgDeleteEvent.deleteEventReq)
	default:
		fmt.Println("bad route")
	}
}

func sendCreateEventMessage(ctx context.Context, cc *grpc.ClientConn, message proto.CreateEventRequestMessage) *proto.CreateEventResponseMessage {
	c := proto.NewCreateEventServiceClient(cc)
	msg, err := c.SendCreateEventMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}
	if msg != nil {
		fmt.Printf("\nerror:%v status:%v\n, record: %#v, id %v\n", msg.Error, msg.Status, msg.Record, msg.Record.Id)
	}

	return msg
}

func sendDeleteEventMessage(ctx context.Context, cc *grpc.ClientConn, message proto.DeleteEventRequestMessage) *proto.DeleteEventResponseMessage {
	c := proto.NewCreateEventServiceClient(cc)
	msg, err := c.SendDeleteEventMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}
	if msg != nil {
		fmt.Printf("\nstatus:%v text:%v\n", msg.Status, msg.Text)
	}

	return msg
}

type Dummy struct {
	createEventReq proto.CreateEventRequestMessage
	deleteEventReq proto.DeleteEventRequestMessage
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.

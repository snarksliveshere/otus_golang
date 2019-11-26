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
	cc, err := grpc.Dial("0.0.0.0:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func() { _ = cc.Close() }()

	msgCreateEvent := Dummy{createEventReq: proto.CreateEventRequestMessage{
		Title:       "Some_title1",
		Description: "Some_description1",
		Date:        "2019-11-01",
	}}
	msgCreateEvent2 := Dummy{createEventReq: proto.CreateEventRequestMessage{
		Title:       "Some_title2",
		Description: "Some_description2",
		Date:        "2019-11-01",
	}}
	msgCreateEvent3 := Dummy{createEventReq: proto.CreateEventRequestMessage{
		Title:       "Some_title3",
		Description: "Some_description3",
		Date:        "2019-11-02",
	}}
	msgCreateEvent4 := Dummy{createEventReq: proto.CreateEventRequestMessage{
		Title:       "Some_title4",
		Description: "Some_description4",
		Date:        "2019-10-02",
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
	case "update-event":
		rec := sendCreateEventMessage(ctx, cc, msgCreateEvent.createEventReq)
		msgUpdateEvent := Dummy{
			updateEventReq: proto.UpdateEventRequestMessage{
				EventId:     rec.Record.Id,
				Title:       "update_title",
				Description: "update_description",
				Date:        "2019-11-01",
			},
		}
		sendUpdateEventMessage(ctx, cc, msgUpdateEvent.updateEventReq)

	case "get-day-events":
		msgDayEvent := Dummy{
			eventForDayReq: proto.GetEventsForDateRequestMessage{
				Date: "2019-11-01",
			},
		}

		sendCreateEventMessage(ctx, cc, msgCreateEvent.createEventReq)
		sendCreateEventMessage(ctx, cc, msgCreateEvent2.createEventReq)
		sendCreateEventMessage(ctx, cc, msgCreateEvent3.createEventReq)
		sendCreateEventMessage(ctx, cc, msgCreateEvent4.createEventReq)
		sendGetEventsForDayMessage(ctx, cc, msgDayEvent.eventForDayReq)
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

func sendUpdateEventMessage(ctx context.Context, cc *grpc.ClientConn, message proto.UpdateEventRequestMessage) *proto.UpdateEventResponseMessage {
	c := proto.NewCreateEventServiceClient(cc)
	msg, err := c.SendUpdateEventMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}
	if msg != nil {
		fmt.Printf("\nstatus:%v text:%v\n", msg.Status, msg.Text)
	}

	return msg
}

func sendGetEventsForDayMessage(ctx context.Context, cc *grpc.ClientConn, message proto.GetEventsForDateRequestMessage) *proto.GetEventsForDateResponseMessage {
	c := proto.NewCreateEventServiceClient(cc)
	msg, err := c.SendGetEventsForDayMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}

	if msg != nil {
		fmt.Printf("\nstatus:%v text:%v, records: %#v, records title1: %#v, records title2: %#v, lengh (must be 2): %d\n",
			msg.Status, msg.Text, msg.Records, msg.Records[0].Title, msg.Records[2].Title, len(msg.Records))
	}

	return msg
}

type Dummy struct {
	createEventReq proto.CreateEventRequestMessage
	deleteEventReq proto.DeleteEventRequestMessage
	updateEventReq proto.UpdateEventRequestMessage
	eventForDayReq proto.GetEventsForDateRequestMessage
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.

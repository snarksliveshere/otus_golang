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
			EventId: rec.Event.Id,
		}}
		sendDeleteEventMessage(ctx, cc, msgDeleteEvent.deleteEventReq)
	case "update-event":
		rec := sendCreateEventMessage(ctx, cc, msgCreateEvent.createEventReq)
		msgUpdateEvent := Dummy{
			updateEventReq: proto.UpdateEventRequestMessage{
				EventId:     rec.Event.Id,
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
	case "get-month-event":
		sendCreateEventMessage(ctx, cc, msgCreateEvent.createEventReq)
		sendCreateEventMessage(ctx, cc, msgCreateEvent2.createEventReq)
		sendCreateEventMessage(ctx, cc, msgCreateEvent4.createEventReq)
		msgMonthEvent := Dummy{
			eventForMonthReq: proto.GetEventsForMonthRequestMessage{
				Month: "2019-11",
			},
		}
		sendGetEventsForMonthMessage(ctx, cc, msgMonthEvent.eventForMonthReq)

	case "get-interval-event":
		sendCreateEventMessage(ctx, cc, msgCreateEvent.createEventReq)
		sendCreateEventMessage(ctx, cc, msgCreateEvent2.createEventReq)
		sendCreateEventMessage(ctx, cc, msgCreateEvent3.createEventReq)
		sendCreateEventMessage(ctx, cc, msgCreateEvent4.createEventReq)
		msgIntervalEvent := Dummy{
			eventForIntervalReq: proto.GetEventsForIntervalRequestMessage{
				From: "2019-10-20",
				Till: "2019-11-01",
			},
		}
		sendGetEventsForIntervalMessage(ctx, cc, msgIntervalEvent.eventForIntervalReq)
	default:
		fmt.Println("bad route")
	}
}

func sendCreateEventMessage(ctx context.Context, cc *grpc.ClientConn, message proto.CreateEventRequestMessage) *proto.CreateEventResponseMessage {
	c := proto.NewEventServiceClient(cc)
	msg, err := c.SendCreateEventMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}
	if msg != nil {
		fmt.Printf("\nerror:%v status:%v\n, event: %#v, id %v\n", msg.Error, msg.Status, msg.Event, msg.Event.Id)
	}

	return msg
}

func sendDeleteEventMessage(ctx context.Context, cc *grpc.ClientConn, message proto.DeleteEventRequestMessage) *proto.DeleteEventResponseMessage {
	c := proto.NewEventServiceClient(cc)
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
	c := proto.NewEventServiceClient(cc)
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
	c := proto.NewEventServiceClient(cc)
	msg, err := c.SendGetEventsForDayMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}

	if msg != nil {
		fmt.Printf("\nstatus:%v text:%v, events: %#v, events title1: %#v, events title2: %#v, date: %v\n",
			msg.Status, msg.Text, msg.Events, msg.Events[0].Title, msg.Events[1].Title, msg.Date)
	}

	return msg
}

func sendGetEventsForMonthMessage(ctx context.Context, cc *grpc.ClientConn, message proto.GetEventsForMonthRequestMessage) *proto.GetEventsForMonthResponseMessage {
	c := proto.NewEventServiceClient(cc)
	msg, err := c.SendGetEventsForMonthMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}

	if msg != nil {
		fmt.Printf("\nstatus:%v text:%v, events: %#v, events title1: %#v, events title2: %#v\n",
			msg.Status, msg.Text, msg.Events, msg.Events[0].Title, msg.Events[1].Title)
	}

	return msg
}

func sendGetEventsForIntervalMessage(ctx context.Context, cc *grpc.ClientConn, message proto.GetEventsForIntervalRequestMessage) *proto.GetEventsForIntervalResponseMessage {
	c := proto.NewEventServiceClient(cc)
	msg, err := c.SendGetEventsForIntervalMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}

	if msg != nil {
		fmt.Printf("\nstatus:%v text:%v, events: %#v, events title1: %#v, events title2: %#v\n",
			msg.Status, msg.Text, msg.Events, msg.Events[0].Title, msg.Events[1].Title)
	}

	return msg
}

type Dummy struct {
	createEventReq      proto.CreateEventRequestMessage
	deleteEventReq      proto.DeleteEventRequestMessage
	updateEventReq      proto.UpdateEventRequestMessage
	eventForDayReq      proto.GetEventsForDateRequestMessage
	eventForMonthReq    proto.GetEventsForMonthRequestMessage
	eventForIntervalReq proto.GetEventsForIntervalRequestMessage
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.

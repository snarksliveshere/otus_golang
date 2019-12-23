package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/client/config"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"time"
)

func createTimeStampFromTimeString(timeStr string) (*timestamp.Timestamp, error) {
	t, err := time.Parse(config.EventTimeLayout, timeStr)
	if err != nil {
		fmt.Println("bad time format")
		return nil, err
	}
	return ptypes.TimestampProto(t)
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	cc, err := grpc.Dial("0.0.0.0:"+config.ConfigPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func() { _ = cc.Close() }()
	time1, err := createTimeStampFromTimeString("2019-05-10T20:03+0300")
	timeDelete, err := createTimeStampFromTimeString("2019-09-02T14:03+0300")
	if err != nil {
		fmt.Println("cant convert time to proto timestamp")
	}

	msgCreateEvent := Dummy{createEventReq: proto.CreateEventRequestMessage{
		Title:       "Some_title1",
		Description: "Some_description1",
		Time:        time1,
	}}
	msgCreateEventDelete := Dummy{createEventReq: proto.CreateEventRequestMessage{
		Title:       "Some_title delete",
		Description: "Some_description delete",
		Time:        timeDelete,
	}}

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		return
	}

	switch expr := os.Args[1]; expr {
	case "get-event-by-id":
		rec := sendCreateEventMessage(ctx, cc, msgCreateEvent.createEventReq)
		msgGetByIdEventDelete := Dummy{eventById: proto.GetEventByIdRequestMessage{
			EventId: rec.Id,
		}}
		sendGetEventsByIdMessage(ctx, cc, msgGetByIdEventDelete.eventById)
	case "create-event":
		sendCreateEventMessage(ctx, cc, msgCreateEvent.createEventReq)
	case "delete-event":
		rec := sendCreateEventMessage(ctx, cc, msgCreateEventDelete.createEventReq)
		msgDeleteEvent := Dummy{deleteEventReq: proto.DeleteEventRequestMessage{
			EventId: rec.Id,
		}}
		sendDeleteEventMessage(ctx, cc, msgDeleteEvent.deleteEventReq)
	case "update-event":
		rec := sendCreateEventMessage(ctx, cc, msgCreateEvent.createEventReq)
		msgUpdateEvent := Dummy{
			updateEventReq: proto.UpdateEventRequestMessage{
				EventId:     rec.Id,
				Title:       "update_title11",
				Description: "update_description11",
			},
		}
		sendUpdateEventMessage(ctx, cc, msgUpdateEvent.updateEventReq)

	case "get-day-events":
		msgDayEvent := Dummy{
			eventForDayReq: proto.GetEventsForDateRequestMessage{
				Date: "2019-11-10",
			},
		}
		sendGetEventsForDayMessage(ctx, cc, msgDayEvent.eventForDayReq)
	case "get-month-event":
		msgMonthEvent := Dummy{
			eventForMonthReq: proto.GetEventsForMonthRequestMessage{
				Month: "2019-11",
			},
		}
		sendGetEventsForMonthMessage(ctx, cc, msgMonthEvent.eventForMonthReq)

	case "get-interval-event":
		msgIntervalEvent := Dummy{
			eventForIntervalReq: proto.GetEventsForIntervalRequestMessage{
				From: "2019-05-09",
				Till: "2019-10-20",
			},
		}
		sendGetEventsForIntervalMessage(ctx, cc, msgIntervalEvent.eventForIntervalReq)
	default:
		fmt.Println("bad route")
	}
}

func sendGetEventsByIdMessage(ctx context.Context, cc *grpc.ClientConn, message proto.GetEventByIdRequestMessage) *proto.GetEventByIdResponseMessage {
	c := proto.NewEventServiceClient(cc)
	msg, err := c.SendGetEventsByIdMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}

	if msg != nil {
		fmt.Printf("\nstatus:%v error:%v, events title1: %#v,  event time: %#v",
			msg.Status, msg.Error, msg.Event.Title, msg.Event.Time)
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
		fmt.Printf("\nstatus:%v text:%v, events: %#v",
			msg.Status, msg.Text, msg.Events)
	}

	return msg
}

func sendCreateEventMessage(ctx context.Context, cc *grpc.ClientConn, message proto.CreateEventRequestMessage) *proto.CreateEventResponseMessage {
	c := proto.NewEventServiceClient(cc)
	msg, err := c.SendCreateEventMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}
	if msg != nil {
		fmt.Printf("\nerror:%v status:%v\n, id %v\n", msg.Error, msg.Status, msg.Id)
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

func sendGetEventsForMonthMessage(ctx context.Context, cc *grpc.ClientConn, message proto.GetEventsForMonthRequestMessage) *proto.GetEventsForMonthResponseMessage {
	c := proto.NewEventServiceClient(cc)
	msg, err := c.SendGetEventsForMonthMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}

	if msg != nil {
		fmt.Printf("\nstatus:%v text:%v, events: %#v, length %v",
			msg.Status, msg.Text, msg.Events, len(msg.Events))
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
		fmt.Printf("\nstatus:%v text:%v, events: %#v, length events: %d\n",
			msg.Status, msg.Text, msg.Events, len(msg.Events))
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
	eventById           proto.GetEventByIdRequestMessage
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.

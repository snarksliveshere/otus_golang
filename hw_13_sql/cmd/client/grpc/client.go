package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/config"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/proto"
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
	cc, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer func() { _ = cc.Close() }()
	time1, err := createTimeStampFromTimeString("2019-11-10T20:03+0300")
	time2, err := createTimeStampFromTimeString("2019-11-02T18:03+0300")
	time3, err := createTimeStampFromTimeString("2019-10-02T14:03+0300")
	if err != nil {
		fmt.Println("cant convert time to proto timestamp")
	}

	msgCreateEvent := Dummy{createEventReq: proto.CreateEventRequestMessage{
		Title:       "Some_title1",
		Description: "Some_description1",
		Time:        time1,
	}}
	msgCreateEvent2 := Dummy{createEventReq: proto.CreateEventRequestMessage{
		Title:       "Some_title2",
		Description: "Some_description2",
		Time:        time2,
	}}
	msgCreateEvent4 := Dummy{createEventReq: proto.CreateEventRequestMessage{
		Title:       "Some_title4",
		Description: "Some_description4",
		Time:        time3,
	}}

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		return
	}

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

func sendGetEventsForDayMessage(ctx context.Context, cc *grpc.ClientConn, message proto.GetEventsForDateRequestMessage) *proto.GetEventsForDateResponseMessage {
	c := proto.NewEventServiceClient(cc)
	msg, err := c.SendGetEventsForDayMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}

	if msg != nil {
		fmt.Printf("\nstatus:%v text:%v, records: %#v, records title1: %#v,  record time: %#v",
			msg.Status, msg.Text, msg.Records, msg.Records[0].Title, msg.Records[0].Time)
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
		fmt.Printf("\nstatus:%v text:%v, records: %#v, records title1: %#v, records title2: %#v\n",
			msg.Status, msg.Text, msg.Records, msg.Records[0].Title, msg.Records[1].Title)
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
		fmt.Printf("\nstatus:%v text:%v, records: %#v, records title1: %#v, records title2: %#v\n",
			msg.Status, msg.Text, msg.Records, msg.Records[0].Title, msg.Records[1].Title)
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

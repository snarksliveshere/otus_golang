package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/client/config"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/client/proto"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var (
	conf config.AppConfig
	cc   *grpc.ClientConn
)

func init() {
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")
	var err error
	cc, err = grpc.Dial(conf.ListenIP+":"+conf.GRPCPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cant connect to host: %s, port: %s, error:%s", conf.ListenIP, conf.GRPCPort, err.Error())
	}
}

func TestSendCreateEventMessage(t *testing.T) {
	cases := []struct {
		status, title, description, time string
	}{
		{
			status:      "success",
			title:       "some new title",
			description: "some new description",
			time:        "2019-12-07T20:03+0300",
		},
		{
			status:      "success",
			title:       "new title2",
			description: "some new description2",
			time:        "2019-12-07T19:03+0300",
		},
		{
			status:      "success",
			title:       "new title3",
			description: "some new description3",
			time:        "2019-12-07T19:30+0300",
		},
		{
			status:      "success",
			title:       "new title3",
			description: "some new description3",
			time:        "2019-12-08T11:33+0300",
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		timeS, err := createTimeStampFromTimeString(c.time)
		if err != nil {
			t.Errorf("TestCreateEvent(),  err time: %s", c.time)
		}
		msg := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Time:        timeS,
		}

		resp := sendCreateEventMessage(ctx, cc, msg)
		msgId := proto.GetEventByIdRequestMessage{
			EventId: resp.Id,
		}
		rec := sendGetEventsByIdMessage(ctx, cc, msgId)

		if c.status != resp.Status {
			t.Errorf("TestSendCreateEventMessage() status compare, c.status: %s, resp.status: %v, resp ID: %d", c.status, resp.Status, resp.Id)
		}

		if rec.Event.Title != c.title && rec.Event.Description != c.description && c.status != resp.Status {
			t.Errorf("TestSendCreateEventMessage() title, description compare, c.title: %s, resp.title: %v, c.description: %v, resp description: %v. resp ID: %d",
				c.title, rec.Event.Title, c.description, rec.Event.Description, rec.Event.Id)
		}
	}
}

func TestSendDeleteEventMessage(t *testing.T) {
	cases := []struct {
		status, title, description, time string
		plus                             uint64
	}{
		{
			status:      "success",
			title:       "some new title",
			description: "some new description",
			time:        "2019-05-01T20:03+0300",
		},
		{
			status:      "error",
			title:       "new title2",
			description: "some new description2",
			time:        "2019-04-02T20:03+0300",
			plus:        1,
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		timeS, err := createTimeStampFromTimeString(c.time)
		if err != nil {
			t.Errorf("TestCreateEvent(),  err time: %s", c.time)
		}
		msg := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Time:        timeS,
		}

		resp := sendCreateEventMessage(ctx, cc, msg)

		delMsg := proto.DeleteEventRequestMessage{
			EventId: resp.Id + c.plus,
		}
		respDelete := sendDeleteEventMessage(ctx, cc, delMsg)

		if c.status != respDelete.Status {
			t.Errorf("TestSendDeleteEventMessage() status compare, c.status: %s, resp.status: %v", c.status, resp.Status)
		}
	}
}

func TestSendUpdateEventMessage(t *testing.T) {
	cases := []struct {
		status, title, description, time string
		updTitle, updDescription         string
	}{
		{
			status:         "success",
			title:          "some new title",
			description:    "some new description",
			time:           "2019-02-01T20:03+0300",
			updTitle:       "update_title",
			updDescription: "update_description",
		},
		{
			status:         "success",
			title:          "new title2",
			description:    "some new description2",
			time:           "2019-03-01T20:03+0300",
			updTitle:       "update_title",
			updDescription: "update_description",
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		timeS, err := createTimeStampFromTimeString(c.time)
		if err != nil {
			t.Errorf("TestCreateEvent(),  err time: %s", c.time)
		}
		msg := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Time:        timeS,
		}

		resp := sendCreateEventMessage(ctx, cc, msg)

		updMsg := proto.UpdateEventRequestMessage{
			EventId:     resp.Id,
			Title:       c.updTitle,
			Description: c.updDescription,
		}
		respUpdate := sendUpdateEventMessage(ctx, cc, updMsg)

		msgId := proto.GetEventByIdRequestMessage{
			EventId: resp.Id,
		}
		rec := sendGetEventsByIdMessage(ctx, cc, msgId)

		if c.status != respUpdate.Status {
			t.Errorf("TestSendUpdateEventMessage() status compare, c.status: %s, resp.status: %v", c.status, resp.Status)
		}
		if c.updTitle != rec.Event.Title && c.status != respUpdate.Status {
			t.Errorf("TestSendUpdateEventMessage() title compare, c.title: %s, rec title: %v",
				c.title, rec.Event.Title)
		}
		if c.updDescription != rec.Event.Description && c.status != respUpdate.Status {
			t.Errorf("TestSendUpdateEventMessage() description compare, c.description: %s, rec description: %v",
				c.description, rec.Event.Description)
		}

	}
}

func TestSendGetEventsForDayMessage(t *testing.T) {
	cases := []struct {
		status, title, description, time, date string
	}{
		{
			status:      "success",
			title:       "some new title",
			description: "some new description",
			time:        "2018-05-02T20:03+0300",
			date:        "2018-05-02",
		},
		{
			status:      "error",
			title:       "some new title2",
			description: "some new description2",
			time:        "2018-05-03T18:03+0300",
			date:        "2018-05-04",
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		timeS, err := createTimeStampFromTimeString(c.time)
		if err != nil {
			t.Errorf("TestCreateEvent(),  err time: %s", c.time)
		}
		msg := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Time:        timeS,
		}

		resp := sendCreateEventMessage(ctx, cc, msg)

		dateMsg := proto.GetEventsForDateRequestMessage{
			Date: c.date,
		}
		respEvents := sendGetEventsForDayMessage(ctx, cc, dateMsg)
		if c.status != respEvents.Status {
			t.Errorf("sendGetEventsForDayMessage() status compare, c.status: %s, resp.status: %v", c.status, resp.Status)
		}
	}
}

func TestSendGetEventsForMonthMessage(t *testing.T) {
	cases := []struct {
		status, title, description, time, time2, date string
		length                                        int
	}{
		{
			status:      "success",
			title:       "some new title",
			description: "some new description",
			time:        "1984-01-02T20:03+0300",
			time2:       "1984-01-03T20:03+0300",
			date:        "1984-01",
			length:      2,
		},
		{
			status:      "error",
			title:       "some new title2",
			description: "some new description2",
			time:        "1985-01-01T20:03+0300",
			time2:       "1985-01-03T20:03+0300",
			date:        "1985-02",
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		timeS, err := createTimeStampFromTimeString(c.time)
		if err != nil {
			t.Errorf("TestSendGetEventsForMonthMessage(),  err time: %s", c.time)
		}
		timeS2, err := createTimeStampFromTimeString(c.time2)
		if err != nil {
			t.Errorf("TestSendGetEventsForMonthMessage(),  err time: %s", c.time2)
		}
		msg := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Time:        timeS,
		}
		msg2 := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Time:        timeS2,
		}

		sendCreateEventMessage(ctx, cc, msg)
		sendCreateEventMessage(ctx, cc, msg2)

		dateMsg := proto.GetEventsForMonthRequestMessage{
			Month: c.date,
		}
		respEvents := sendGetEventsForMonthMessage(ctx, cc, dateMsg)
		if c.status != respEvents.Status {
			t.Errorf("sendGetEventsForMonthMessage() status compare, c.status: %s, resp.status: %v", c.status, respEvents.Status)
		}
		if c.length != len(respEvents.Events) {
			t.Errorf("sendGetEventsForMonthMessage() length compare, c.length: %d, resp.length: %d", c.length, len(respEvents.Events))
		}
	}
}

func TestSendGetEventsForIntervalMessage(t *testing.T) {
	cases := []struct {
		status, title, description, time, time2, from, till string
		length                                              int
	}{
		{
			status:      "success",
			title:       "some new title",
			description: "some new description",
			time:        "1979-01-02T20:03+0300",
			time2:       "1979-03-03T20:03+0300",
			from:        "1979-01-02",
			till:        "1979-03-04",
			length:      2,
		},
		{
			status:      "error",
			title:       "some new title2",
			description: "some new description2",
			time:        "1968-01-01T20:03+0300",
			time2:       "1968-01-03T20:03+0300",
			from:        "1969-01-02",
			till:        "1969-03-04",
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		timeS, err := createTimeStampFromTimeString(c.time)
		if err != nil {
			t.Errorf("TestSendGetEventsForIntervalMessage(),  err time: %s", c.time)
		}
		timeS2, err := createTimeStampFromTimeString(c.time2)
		if err != nil {
			t.Errorf("TestSendGetEventsForIntervalMessage(),  err time: %s", c.time2)
		}
		msg := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Time:        timeS,
		}
		msg2 := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Time:        timeS2,
		}

		sendCreateEventMessage(ctx, cc, msg)
		sendCreateEventMessage(ctx, cc, msg2)

		dateMsg := proto.GetEventsForIntervalRequestMessage{
			From: c.from,
			Till: c.till,
		}
		respEvents := sendGetEventsForIntervalMessage(ctx, cc, dateMsg)
		if c.status != respEvents.Status {
			t.Errorf("TestSendGetEventsForIntervalMessage() status compare, c.status: %s, resp.status: %v", c.status, respEvents.Status)
		}
		if c.length != len(respEvents.Events) {
			t.Errorf("TestSendGetEventsForIntervalMessage() length compare, c.length: %d, resp.length: %d", c.length, len(respEvents.Events))
		}
	}
}

func TestDummyGetEventForTimeIntervalMessage(t *testing.T) {
	cases := []struct {
		status, title, description, add string
	}{
		{
			status:      "success",
			title:       "some new title",
			description: "some new description",
		},
		{
			status:      "success",
			title:       "new title2",
			description: "some new description2",
			add:         "+8",
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		from, till, err := getTimestampsInterval()
		if err != nil {
			t.Errorf("TestDummyCreateEvent(), resp.status: %s", c.status)
		}
		msg := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
		}

		if c.add == "" {
			msg.Time = from
		} else {
			msg.Time = till
		}

		resp := sendCreateEventMessage(ctx, cc, msg)
		msgId := proto.GetEventByIdRequestMessage{
			EventId: resp.Id,
		}
		sendGetEventsByIdMessage(ctx, cc, msgId)

		if c.status != resp.Status {
			t.Errorf("TestSendCreateEventMessage() status compare, c.status: %s, resp.status: %v, resp ID: %d", c.status, resp.Status, resp.Id)
		}
	}
}

func getTimestampsInterval() (*timestamp.Timestamp, *timestamp.Timestamp, error) {
	from := time.Now()
	till := from.Add(8 * time.Minute)
	fromT, err := ptypes.TimestampProto(from)
	if err != nil {
		return nil, nil, err
	}
	tillT, err := ptypes.TimestampProto(till)
	if err != nil {
		return nil, nil, err
	}
	return fromT, tillT, nil
}

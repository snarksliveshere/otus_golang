package main

import (
	"context"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/config"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/proto"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestSendCreateEventMessage(t *testing.T) {
	cases := []struct {
		status, title, description, time string
	}{
		{
			status:      "success",
			title:       "some new title",
			description: "some new description",
			time:        "2019-05-02T20:03+0300",
		},
		{
			status:      "success",
			title:       "new title2",
			description: "some new description2",
			time:        "2019-06-10T20:03+0300",
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		cc, err := grpc.Dial("0.0.0.0:"+config.ConfigPort, grpc.WithInsecure())
		if err != nil {
			t.Errorf("TestCreateEvent(), resp.status: %s", c.status)
		}
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

		if rec.Record.Title != c.title && rec.Record.Description != c.description && c.status != resp.Status {
			t.Errorf("TestSendCreateEventMessage() title, description compare, c.title: %s, resp.title: %v, c.description: %v, resp description: %v. resp ID: %d",
				c.title, rec.Record.Title, c.description, rec.Record.Description, rec.Record.Id)
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
		cc, err := grpc.Dial("0.0.0.0:"+config.ConfigPort, grpc.WithInsecure())
		if err != nil {
			t.Errorf("TestCreateEvent(), resp.status: %s", c.status)
		}
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
		plus                             uint64
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
			status:         "error",
			title:          "new title2",
			description:    "some new description2",
			time:           "2019-03-01T20:03+0300",
			plus:           1,
			updTitle:       "update_title",
			updDescription: "update_description",
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		cc, err := grpc.Dial("0.0.0.0:"+config.ConfigPort, grpc.WithInsecure())
		if err != nil {
			t.Errorf("TestCreateEvent(), resp.status: %s", c.status)
		}
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
			EventId:     resp.Id + c.plus,
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
		if c.updTitle != rec.Record.Title && c.status != respUpdate.Status {
			t.Errorf("TestSendUpdateEventMessage() title compare, c.title: %s, rec title: %v",
				c.title, rec.Record.Title)
		}
		if c.updDescription != rec.Record.Description && c.status != respUpdate.Status {
			t.Errorf("TestSendUpdateEventMessage() description compare, c.description: %s, rec description: %v",
				c.description, rec.Record.Description)
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
		cc, err := grpc.Dial("0.0.0.0:"+config.ConfigPort, grpc.WithInsecure())
		if err != nil {
			t.Errorf("TestCreateEvent(), resp.status: %s", c.status)
		}
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
		respRecords := sendGetEventsForDayMessage(ctx, cc, dateMsg)
		if c.status != respRecords.Status {
			t.Errorf("sendGetEventsForDayMessage() status compare, c.status: %s, resp.status: %v", c.status, resp.Status)
		}
	}
}

package main

import (
	"context"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/config"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/proto"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestSendCreateEventMessage(t *testing.T) {
	cases := []struct {
		status, title, description, date string
	}{
		{
			status:      "success",
			title:       "some new title",
			description: "some new description",
			date:        "2019-11-01",
		},
		{
			status:      "success",
			title:       "new title2",
			description: "some new description2",
			date:        "2019-11-01",
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		cc, err := grpc.Dial("0.0.0.0:"+config.ConfigPort, grpc.WithInsecure())
		if err != nil {
			t.Errorf("TestCreateEvent(), resp.status: %s", c.status)
		}
		msg := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Date:        c.date,
		}

		resp := sendCreateEventMessage(ctx, cc, msg)

		if c.status != resp.Status {
			t.Errorf("TestSendCreateEventMessage() status compare, c.status: %s, resp.status: %v", c.status, resp.Status)
		}

		if c.title != resp.Record.Title {
			t.Errorf("TestSendCreateEventMessage() title compare, c.title: %s, resp.title: %v", c.title, resp.Record.Title)
		}

		if c.description != resp.Record.Description {
			t.Errorf("TestSendCreateEventMessage() description compare, c.description: %s, resp.description: %v", c.description, resp.Record.Description)
		}

		if resp.Record.Id <= 0 {
			t.Errorf("TestSendCreateEventMessage() must be id, resp id: %v", resp.Record.Id)
		}
	}
}

func TestSendDeleteEventMessage(t *testing.T) {
	cases := []struct {
		status, title, description, date string
		plus                             uint64
	}{
		{
			status:      "success",
			title:       "some new title",
			description: "some new description",
			date:        "2019-11-01",
		},
		{
			status:      "error",
			title:       "new title2",
			description: "some new description2",
			date:        "2019-11-01",
			plus:        1,
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		cc, err := grpc.Dial("0.0.0.0:"+config.ConfigPort, grpc.WithInsecure())
		if err != nil {
			t.Errorf("TestCreateEvent(), resp.status: %s", c.status)
		}
		msg := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Date:        c.date,
		}

		resp := sendCreateEventMessage(ctx, cc, msg)

		delMsg := proto.DeleteEventRequestMessage{
			EventId: resp.Record.Id + c.plus,
		}
		respDelete := sendDeleteEventMessage(ctx, cc, delMsg)

		if c.status != respDelete.Status {
			t.Errorf("TestSendDeleteEventMessage() status compare, c.status: %s, resp.status: %v", c.status, resp.Status)
		}
	}
}

func TestSendUpdateEventMessage(t *testing.T) {
	cases := []struct {
		status, title, description, date string
		updTitle, updDescription         string
		plus                             uint64
	}{
		{
			status:         "success",
			title:          "some new title",
			description:    "some new description",
			date:           "2019-11-01",
			updTitle:       "update_title",
			updDescription: "update_description",
		},
		{
			status:         "error",
			title:          "new title2",
			description:    "some new description2",
			date:           "2019-11-01",
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
		msg := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Date:        c.date,
		}

		resp := sendCreateEventMessage(ctx, cc, msg)

		updMsg := proto.UpdateEventRequestMessage{
			EventId:     resp.Record.Id + c.plus,
			Title:       c.updTitle,
			Description: c.updDescription,
			Date:        c.date,
		}
		respUpdate := sendUpdateEventMessage(ctx, cc, updMsg)

		if c.status != respUpdate.Status {
			t.Errorf("TestSendUpdateEventMessage() status compare, c.status: %s, resp.status: %v", c.status, resp.Status)
		}
	}
}

func TestSendGetEventsForDayMessage(t *testing.T) {
	cases := []struct {
		status, title, description, date string
	}{
		{
			status:      "success",
			title:       "some new title",
			description: "some new description",
			date:        "2019-11-01",
		},
		{
			status:      "success",
			title:       "some new title2",
			description: "some new description2",
			date:        "2019-11-01",
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		cc, err := grpc.Dial("0.0.0.0:"+config.ConfigPort, grpc.WithInsecure())
		if err != nil {
			t.Errorf("TestCreateEvent(), resp.status: %s", c.status)
		}
		msg := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Date:        c.date,
		}

		resp := sendCreateEventMessage(ctx, cc, msg)

		dateMsg := proto.GetEventsForDateRequestMessage{
			Date: c.date,
		}
		respRecords := sendGetEventsForDayMessage(ctx, cc, dateMsg)
		if c.date != respRecords.Date {
			t.Errorf("sendGetEventsForDayMessage() status compare, c.status: %s, resp.status: %v", c.status, resp.Status)
		}
	}
}

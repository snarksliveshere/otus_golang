package main

import (
	"context"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/proto"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestCreateEvent(t *testing.T) {
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
		cc, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
		if err != nil {
			t.Errorf("TestCreateEvent(), resp.status: %s", c.status)
		}
		msg := proto.CreateEventRequestMessage{
			Title:       c.title,
			Description: c.description,
			Date:        c.date,
		}

		resp := sendCreateMessage(ctx, cc, msg)

		if c.status != resp.Status {
			t.Errorf("TestCreateEvent() status compare, c.status: %s, resp.status: %v", c.status, resp.Status)
		}

		if c.title != resp.Record.Title {
			t.Errorf("TestCreateEvent() title compare, c.title: %s, resp.title: %v", c.title, resp.Record.Title)
		}

		if c.description != resp.Record.Description {
			t.Errorf("TestCreateEvent() description compare, c.description: %s, resp.description: %v", c.description, resp.Record.Description)
		}

		if resp.Record.Id <= 0 {
			t.Errorf("TestCreateEvent() must be id, resp id: %v", resp.Record.Id)
		}
	}
}

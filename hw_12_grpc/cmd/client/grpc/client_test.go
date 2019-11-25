package main

import (
	"context"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestCreateEvent(t *testing.T) {
	cases := []struct {
		status string
	}{
		{
			status: "success",
		},
	}

	for _, c := range cases {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		cc, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
		if err != nil {
			t.Errorf("TestCreateEvent(), resp.status: %s", c.status)
		}
		resp := sendCreateMessage(ctx, cc)

		if c.status != resp.Status {
			t.Errorf("TestCreateEvent(), resp.status: %s", c.status)
		}
	}
}

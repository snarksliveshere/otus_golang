package grpc

import (
	"context"
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/entity"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/internal/data_handlers"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Response struct {
	Date    entity.Date     `json:"day,omitempty"`
	Record  entity.Record   `json:"record,omitempty"`
	Records []entity.Record `json:"records,omitempty"`
	//Result     []string      `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
	Status string `json:"status,omitempty"`
}

func (s ServerCalendar) SendCreateEventMessage(ctx context.Context, msg *proto.CreateEventRequestMessage) (*proto.CreateEventResponseMessage, error) {
	title, desc, day, err := data_handlers.CheckCreateEvent(msg.Title, msg.Description, msg.Date)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid title, desc string")
	}
	rec, dt, c, err := storage.AddRecord(title, desc, day)
	reply := proto.CreateEventResponseMessage{}

	if err != nil {
		reply.Status = "error"
		reply.Error = err.Error()
	}
	fmt.Println(rec, dt, c)

	reply.Status = "success"
	reply.Record = entity.Record{
		Id:          1,
		Title:       "title",
		Description: "description",
	}

	return &reply, nil

}

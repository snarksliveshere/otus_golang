package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/entity"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/internal/data_handlers"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
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

	recBytes, err := json.Marshal(&rec)

	fmt.Println(string(recBytes))

	mr := &proto.Record{}
	rN := strings.NewReader(string(recBytes))

	if err := jsonpb.Unmarshal(rN, mr); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//err = mr.XXX_Unmarshal(recBytes)
	//if err!= nil {
	//	fmt.Println("ola")
	//	return nil, status.Error(codes.InvalidArgument, err.Error())
	//}

	reply.Status = "success"

	reply.Record = mr
	//reply.Record = &proto.Record{
	//	Id:          rec.Id,
	//	Title:       rec.Title,
	//	Description: rec.Description,
	//}

	return &reply, nil

}

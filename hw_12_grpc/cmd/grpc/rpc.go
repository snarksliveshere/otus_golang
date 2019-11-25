package grpc

import (
	"context"
	"encoding/json"
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
	rec, _, _, err := storage.AddRecord(title, desc, day)
	reply := proto.CreateEventResponseMessage{}

	if err != nil {
		reply.Status = "error"
		reply.Error = err.Error()
	}

	protoRecord, err := recordToProtoStruct(&rec)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	//recBytes, err := json.Marshal(&rec)
	//if err != nil {
	//	return nil, status.Error(codes.Aborted, err.Error())
	//}
	//
	//protoRecord := &proto.Record{}
	//recordBytesReader := strings.NewReader(string(recBytes))
	//
	//if err := jsonpb.Unmarshal(recordBytesReader, protoRecord); err != nil {
	//	return nil, status.Error(codes.Aborted, err.Error())
	//}
	reply.Status = "success"

	reply.Record = protoRecord

	return &reply, nil
}

func recordToProtoStruct(record *entity.Record) (*proto.Record, error) {
	recBytes, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}
	protoRecord := &proto.Record{}
	recordBytesReader := strings.NewReader(string(recBytes))

	if err := jsonpb.Unmarshal(recordBytesReader, protoRecord); err != nil {
		return nil, err
	}

	return protoRecord, nil
}

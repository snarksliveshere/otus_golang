package grpc

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/config"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/data_handlers"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type Response struct {
	Date    entity.Date     `json:"day,omitempty"`
	Record  entity.Record   `json:"record,omitempty"`
	Records []entity.Record `json:"records,omitempty"`
	Error   string          `json:"error,omitempty"`
	Status  string          `json:"status,omitempty"`
	//Result     []string      `json:"result,omitempty"`
}

func (s ServerCalendar) SendCreateEventMessage(ctx context.Context, msg *proto.CreateEventRequestMessage) (*proto.CreateEventResponseMessage, error) {
	title, desc, time, err := data_handlers.CheckCreateEventProtoTimestamp(msg.Title, msg.Description, msg.Time)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid title, desc string")
	}

	id, err := storage.Actions.CreateEvent(title, desc, time)
	reply := proto.CreateEventResponseMessage{}

	if err != nil {
		reply.Status = config.StatusError
		reply.Error = err.Error()
		return &reply, nil
	}
	reply.Status = config.StatusSuccess
	reply.Id = id

	return &reply, nil
}

func (s ServerCalendar) SendGetEventsForDayMessage(ctx context.Context, msg *proto.GetEventsForDateRequestMessage) (*proto.GetEventsForDateResponseMessage, error) {
	records, err := storage.Actions.GetEventsByDay(msg.Date)
	reply := proto.GetEventsForDateResponseMessage{}

	if err != nil {
		reply.Status = config.StatusError
		reply.Text = err.Error()
		return &reply, nil
	}

	var protoRecords []*proto.Record
	for _, record := range records {
		protoRecord, err := recordToProtoStruct(&record)
		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		protoRecords = append(protoRecords, protoRecord)
	}

	reply.Status = config.StatusSuccess
	reply.Records = protoRecords

	return &reply, nil
}

func (s ServerCalendar) SendDeleteEventMessage(ctx context.Context, msg *proto.DeleteEventRequestMessage) (*proto.DeleteEventResponseMessage, error) {
	err := storage.Actions.DeleteRecordById(msg.EventId)
	reply := proto.DeleteEventResponseMessage{}

	if err != nil {
		reply.Status = config.StatusError
		reply.Text = err.Error()
		return &reply, nil
	}
	reply.Status = config.StatusSuccess

	return &reply, nil
}

func (s ServerCalendar) SendUpdateEventMessage(ctx context.Context, msg *proto.UpdateEventRequestMessage) (*proto.UpdateEventResponseMessage, error) {
	title, desc, err := data_handlers.CheckUpdateEventWithoutEventId(msg.Title, msg.Description)
	if err != nil {
		return nil, status.Error(codes.Aborted, "invalid title, desc, date string")
	}
	err = storage.Actions.UpdateRecordById(msg.EventId, title, desc)

	reply := proto.UpdateEventResponseMessage{}
	if err != nil {
		reply.Status = config.StatusError
		reply.Text = err.Error()
		return &reply, nil
	}
	reply.Status = config.StatusSuccess

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

func (s ServerCalendar) SendGetEventsForMonthMessage(ctx context.Context, msg *proto.GetEventsForMonthRequestMessage) (*proto.GetEventsForMonthResponseMessage, error) {
	dates, err := data_handlers.CheckEventsForMonthString(msg.Month)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid month string")
	}
	records, err := storage.Actions.RecordRepository.GetEventsByDateInterval(dates["firstDate"], dates["lastDate"])

	reply := proto.GetEventsForMonthResponseMessage{}

	if err != nil {
		reply.Status = config.StatusError
		reply.Text = err.Error()
		return &reply, nil
	}

	var protoRecords []*proto.Record
	for _, rec := range records {
		protoRecord, err := recordToProtoStruct(&rec)
		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		protoRecords = append(protoRecords, protoRecord)
	}

	reply.Status = config.StatusSuccess
	reply.Records = protoRecords

	return &reply, nil
}

func (s ServerCalendar) SendGetEventsForIntervalMessage(ctx context.Context, msg *proto.GetEventsForIntervalRequestMessage) (*proto.GetEventsForIntervalResponseMessage, error) {
	records, err := storage.Actions.RecordRepository.GetEventsByDateInterval(msg.From, msg.Till)

	reply := proto.GetEventsForIntervalResponseMessage{}

	if err != nil {
		reply.Status = config.StatusError
		reply.Text = err.Error()
		return &reply, nil
	}

	var protoRecords []*proto.Record
	for _, rec := range records {
		protoRecord, err := recordToProtoStruct(&rec)
		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		protoRecords = append(protoRecords, protoRecord)
	}

	reply.Status = config.StatusSuccess
	reply.Records = protoRecords

	return &reply, nil
}

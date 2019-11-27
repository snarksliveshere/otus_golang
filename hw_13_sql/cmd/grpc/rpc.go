package grpc

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/config"
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
		reply.Status = config.StatusError
		reply.Error = err.Error()
		return &reply, nil
	}

	protoRecord, err := recordToProtoStruct(&rec)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	reply.Status = config.StatusSuccess
	reply.Record = protoRecord

	return &reply, nil
}

func (s ServerCalendar) SendDeleteEventMessage(ctx context.Context, msg *proto.DeleteEventRequestMessage) (*proto.DeleteEventResponseMessage, error) {
	eventId := msg.EventId
	err := storage.DeleteRecordById(eventId)
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
	title, desc, day, err := data_handlers.CheckUpdateEventWithoutEventId(msg.Title, msg.Description, msg.Date)
	if err != nil {
		return nil, status.Error(codes.Aborted, "invalid title, desc, date string")
	}
	reply := proto.UpdateEventResponseMessage{}
	err = storage.UpdateRecordById(msg.EventId, day, title, desc)

	if err != nil {
		reply.Status = config.StatusError
		reply.Text = err.Error()
		return &reply, nil
	}
	reply.Status = config.StatusSuccess

	return &reply, nil
}

func (s ServerCalendar) SendGetEventsForDayMessage(ctx context.Context, msg *proto.GetEventsForDateRequestMessage) (*proto.GetEventsForDateResponseMessage, error) {
	t, err := data_handlers.CheckEventsForDay(msg.Date)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid date string")
	}
	day, err := storage.GetEventsForDay(t)

	reply := proto.GetEventsForDateResponseMessage{}

	if err != nil {
		reply.Status = config.StatusError
		reply.Text = err.Error()
		return &reply, nil
	}

	var protoRecords []*proto.Record
	for _, rec := range day.Records {
		protoRecord, err := recordToProtoStruct(&rec)
		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		protoRecords = append(protoRecords, protoRecord)
	}

	reply.Status = config.StatusSuccess
	reply.Records = protoRecords
	reply.Date = day.Day.Format(config.TimeLayout)

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
	dates, err := data_handlers.CheckEventsForMonth(msg.Month)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid month string")
	}
	records, err := storage.GetEventsForInterval(dates["firstDate"], dates["lastDate"])

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
	tFrom, tTill, err := data_handlers.CheckEventsForInterval(msg.From, msg.Till)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid date from date till arguments")
	}
	records, err := storage.GetEventsForInterval(tFrom, tTill)

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
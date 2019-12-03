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
	Date   entity.Date    `json:"day,omitempty"`
	Event  entity.Event   `json:"event,omitempty"`
	Events []entity.Event `json:"events,omitempty"`
	//Result     []string      `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
	Status string `json:"status,omitempty"`
}

func (s ServerCalendar) SendCreateEventMessage(ctx context.Context, msg *proto.CreateEventRequestMessage) (*proto.CreateEventResponseMessage, error) {
	title, desc, day, err := data_handlers.CheckCreateEvent(msg.Title, msg.Description, msg.Date)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid title, desc string")
	}
	rec, _, _, err := storage.AddEvent(title, desc, day)
	reply := proto.CreateEventResponseMessage{}

	if err != nil {
		reply.Status = config.StatusError
		reply.Error = err.Error()
		return &reply, nil
	}

	protoEvent, err := eventToProtoStruct(&rec)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	reply.Status = config.StatusSuccess
	reply.Event = protoEvent

	return &reply, nil
}

func (s ServerCalendar) SendDeleteEventMessage(ctx context.Context, msg *proto.DeleteEventRequestMessage) (*proto.DeleteEventResponseMessage, error) {
	eventId := msg.EventId
	err := storage.DeleteEventById(eventId)
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
	err = storage.UpdateEventById(msg.EventId, day, title, desc)

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

	var protoEvents []*proto.Event
	for _, rec := range day.Events {
		protoEvent, err := eventToProtoStruct(&rec)
		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		protoEvents = append(protoEvents, protoEvent)
	}

	reply.Status = config.StatusSuccess
	reply.Events = protoEvents
	reply.Date = day.Day.Format(config.TimeLayout)

	return &reply, nil
}

func eventToProtoStruct(event *entity.Event) (*proto.Event, error) {
	recBytes, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	protoEvent := &proto.Event{}
	eventBytesReader := strings.NewReader(string(recBytes))

	if err := jsonpb.Unmarshal(eventBytesReader, protoEvent); err != nil {
		return nil, err
	}

	return protoEvent, nil
}

func (s ServerCalendar) SendGetEventsForMonthMessage(ctx context.Context, msg *proto.GetEventsForMonthRequestMessage) (*proto.GetEventsForMonthResponseMessage, error) {
	dates, err := data_handlers.CheckEventsForMonth(msg.Month)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid month string")
	}
	events, err := storage.GetEventsForInterval(dates["firstDate"], dates["lastDate"])

	reply := proto.GetEventsForMonthResponseMessage{}

	if err != nil {
		reply.Status = config.StatusError
		reply.Text = err.Error()
		return &reply, nil
	}

	var protoEvents []*proto.Event
	for _, rec := range events {
		protoEvent, err := eventToProtoStruct(&rec)
		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		protoEvents = append(protoEvents, protoEvent)
	}

	reply.Status = config.StatusSuccess
	reply.Events = protoEvents

	return &reply, nil
}

func (s ServerCalendar) SendGetEventsForIntervalMessage(ctx context.Context, msg *proto.GetEventsForIntervalRequestMessage) (*proto.GetEventsForIntervalResponseMessage, error) {
	tFrom, tTill, err := data_handlers.CheckEventsForInterval(msg.From, msg.Till)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid date from date till arguments")
	}
	events, err := storage.GetEventsForInterval(tFrom, tTill)

	reply := proto.GetEventsForIntervalResponseMessage{}

	if err != nil {
		reply.Status = config.StatusError
		reply.Text = err.Error()
		return &reply, nil
	}

	var protoEvents []*proto.Event
	for _, rec := range events {
		protoEvent, err := eventToProtoStruct(&rec)
		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		protoEvents = append(protoEvents, protoEvent)
	}

	reply.Status = config.StatusSuccess
	reply.Events = protoEvents

	return &reply, nil
}

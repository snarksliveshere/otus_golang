package grpc

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/jsonpb"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/config"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/entity"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/internal/data_handlers"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type Response struct {
	Date   entity.Date    `json:"day,omitempty"`
	Event  entity.Event   `json:"event,omitempty"`
	Events []entity.Event `json:"events,omitempty"`
	Error  string         `json:"error,omitempty"`
	Status string         `json:"status,omitempty"`
	//Result     []string      `json:"result,omitempty"`
}

func (s ServerCalendar) SendGetEventsByIdMessage(ctx context.Context, msg *proto.GetEventByIdRequestMessage) (*proto.GetEventByIdResponseMessage, error) {
	event, err := storage.Actions.EventRepository.FindById(msg.EventId)
	reply := proto.GetEventByIdResponseMessage{}

	if err != nil {
		reply.Status = config.StatusError
		reply.Error = err.Error()
		return &reply, nil
	}
	protoEvent, err := eventToProtoStruct(&event)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}

	reply.Status = config.StatusSuccess
	reply.Event = protoEvent

	return &reply, nil

}
func (s ServerCalendar) SendGetEventsForDayMessage(ctx context.Context, msg *proto.GetEventsForDateRequestMessage) (*proto.GetEventsForDateResponseMessage, error) {
	events, err := storage.Actions.GetEventsByDay(msg.Date)
	reply := proto.GetEventsForDateResponseMessage{}

	if err != nil {
		reply.Status = config.StatusError
		reply.Text = err.Error()
		return &reply, nil
	}

	var protoEvents []*proto.Event
	for _, event := range events {
		protoEvent, err := eventToProtoStruct(&event)
		if err != nil {
			return nil, status.Error(codes.Aborted, err.Error())
		}
		protoEvents = append(protoEvents, protoEvent)
	}

	reply.Status = config.StatusSuccess
	reply.Events = protoEvents

	return &reply, nil
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

func (s ServerCalendar) SendDeleteEventMessage(ctx context.Context, msg *proto.DeleteEventRequestMessage) (*proto.DeleteEventResponseMessage, error) {
	err := storage.Actions.DeleteEventById(msg.EventId)
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
	err = storage.Actions.UpdateEventById(msg.EventId, title, desc)

	reply := proto.UpdateEventResponseMessage{}
	if err != nil {
		reply.Status = config.StatusError
		reply.Text = err.Error()
		return &reply, nil
	}
	reply.Status = config.StatusSuccess

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
	dates, err := data_handlers.CheckEventsForMonthString(msg.Month)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid month string")
	}
	events, err := storage.Actions.EventRepository.GetEventsByDateInterval(dates["firstDate"], dates["lastDate"])

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
	events, err := storage.Actions.EventRepository.GetEventsByDateInterval(msg.From, msg.Till)

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

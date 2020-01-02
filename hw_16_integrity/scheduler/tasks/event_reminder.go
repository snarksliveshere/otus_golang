package tasks

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/scheduler/config"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/scheduler/pkg/logger"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/scheduler/proto"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"time"
)

func EventReminder(log *logger.Logger, ch *amqp.Channel, rk string, conf *config.AppConfig) {
	cc := startGRPCClient(log, conf)
	defer func() { _ = cc.Close() }()
	resp, err := sendGetEventsForTimeIntervalMs(cc)
	if err == nil && resp.Status == config.StatusSuccess {
		startMsg(ch, rk, resp.Events)
	}
}

func getTimestampsInterval() (*timestamp.Timestamp, *timestamp.Timestamp, error) {
	from := time.Now()
	till := from.Add(10 * time.Minute)
	fromT, err := ptypes.TimestampProto(from)
	if err != nil {
		return nil, nil, err
	}
	tillT, err := ptypes.TimestampProto(till)
	if err != nil {
		return nil, nil, err
	}
	return fromT, tillT, nil
}

func startGRPCClient(log *logger.Logger, conf *config.AppConfig) *grpc.ClientConn {
	cc, err := grpc.Dial(conf.GRPCHost+":"+conf.GRPCPort, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	return cc
}

func sendGetEventsForTimeIntervalMs(cc *grpc.ClientConn) (*proto.GetEventsForTimeIntervalResponseMessage, error) {
	from, till, err := getTimestampsInterval()
	if err != nil {
		return nil, err
	}
	message := proto.GetEventsForTimeIntervalRequestMessage{
		From: from,
		Till: till,
	}

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	c := proto.NewEventServiceClient(cc)
	msg, err := c.SendGetEventsForTimeIntervalMessage(ctx, &message)
	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}

	if msg != nil {
		fmt.Printf("\nstatus:%v text:%v, events: %#v",
			msg.Status, msg.Text, msg.Events)
	}
	return msg, nil
}

func startMsg(ch *amqp.Channel, rk, resp string) {
	fmt.Println("start publish!")
	err := ch.Publish(
		"",    // exchange
		rk,    // routing key
		false, // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(resp),
		})
	if err != nil {
		fmt.Println(err.Error())
	}

}

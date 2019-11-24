package main

import (
	"context"
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

//func writeRoutine(end chan interface{}, ctx context.Context, conn chat.ChatExampleClient) {
//	scanner := bufio.NewScanner(os.Stdin)
//OUTER:
//	for {
//		select {
//		case <-ctx.Done():
//			break OUTER
//		default:
//			if !scanner.Scan() {
//				break OUTER
//			}
//			str := scanner.Text()
//			if str == "exit" {
//				break OUTER
//			}
//			msg, err := conn.SendMessage(context.Background(), &chat.ChatMessage{Text: str,})
//			if err != nil {
//				fmt.Printf("error:%s\n", status.Convert(err).Message())
//			}
//
//			if msg != nil {
//				created, _ := ptypes.Timestamp(msg.Created)
//				fmt.Printf("[%s]id:%d msg:%s\n", created.Local(), msg.Id, msg.Text)
//			}
//
//		}
//
//	}
//	log.Printf("Finished writeRoutine")
//	close(end)
//}

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	cc, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := proto.NewCreateEventServiceClient(cc)
	message := proto.CreateEventRequestMessage{
		Title:       "some title",
		Description: "some description",
		Date:        "2019-11-01",
	}
	fmt.Println("find1")
	msg, err := c.SendCreateEventMessage(ctx, &message)

	if err != nil {
		fmt.Printf("error : %s\n", status.Convert(err).Message())
	}

	if msg != nil {
		fmt.Printf("error:%v status:%v\n, record: %#v", msg.Error, msg.Status, msg.Record)
	}
}

//protoc ./proto/events.proto --go_out=plugins=grpc:.

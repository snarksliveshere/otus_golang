package tasks

import (
	"errors"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/scheduler/pkg/database/postgres"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/scheduler/pkg/logger"
	"github.com/streadway/amqp"
	"strconv"
	"time"
)

func EventReminder(log *logger.Logger, db *pg.DB, ch *amqp.Channel, rk string) {
	from := time.Now()
	till := from.Add(10 * time.Minute)
	fmt.Println(from.String(), till.String())
	ids, err := getEventsByDay(from, till, db)
	if err == nil {
		fmt.Printf("\nDONE! []uint64: %#v, err: %#v\n", ids, err)
		idStr := strconv.Itoa(int(ids[0]))
		startMsg(ch, rk, idStr)
	}

	//return getEventsByDay(from, till, db)

}

func startMsg(ch *amqp.Channel, rk, ss string) {
	fmt.Println("start publish!")
	err := ch.Publish(
		"",    // exchange
		rk,    // routing key
		false, // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(ss),
		})
	if err != nil {
		fmt.Println(err.Error())
	}

}

func getEventsByDay(from, till time.Time, db *pg.DB) ([]uint64, error) {
	var loadedRows []*postgres.Event
	err := db.Model(&loadedRows).
		Column("id").
		Where("time >= ?", from).
		Where("time <= ?", till).
		Select()
	if err != nil {
		return nil, err
	}

	var ids []uint64
	for _, v := range loadedRows {

		ids = append(ids, v.Id)
	}

	if len(ids) == 0 {
		return nil, errors.New("there are no events in this day")
	}

	return ids, nil
}

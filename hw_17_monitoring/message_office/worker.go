package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/message_office/config"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/message_office/model"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	msgInSec := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "messages_in_sec",
		})
	prometheus.MustRegister(msgInSec)

	var conf config.AppConfig
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")
	strDial := "amqp://" + conf.RbUser + ":" + conf.RbPassword + "@" + conf.RbHost + ":" + conf.RbPort + "/"
	conn, err := amqp.Dial(strDial)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func() { _ = conn.Close() }()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func() { _ = ch.Close() }()

	q, err := ch.QueueDeclare(
		"events", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		msgInSec.Set(float64(len(msgs)))

		for i := 0; i < 100; i++ {
			msgInSec.Set(float64(i)) // or: Inc(), Dec(), Add(5), Dec(5)
			time.Sleep(10 * time.Millisecond)
		}

		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			err := d.Ack(false)
			if err != nil {
				log.Printf("error: %v", err.Error())
			}
			err = insertToDb(string(d.Body), &conf)
			if err != nil {
				log.Printf("error: %v", err.Error())
			}

		}
	}()
	go func() {
		if err := http.ListenAndServe(":9102", promhttp.Handler()); err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func insertToDb(msg string, conf *config.AppConfig) error {
	db := model.DB{Conf: conf}
	m := model.Message{
		Status: "Success",
		Msg:    msg,
	}
	_, err := db.CreatePgConn().Model(&m).
		Insert()

	if err != nil {
		return err
	}
	return nil
}

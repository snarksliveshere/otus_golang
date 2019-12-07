package main

import (
	"flag"
	"github.com/go-pg/pg"
	"github.com/robfig/cron"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/scheduler/config"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/scheduler/pkg/database/postgres"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/scheduler/pkg/logger"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/scheduler/tasks"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	pathConfig string
)

const (
	confFile = "./config/config.yaml"
)

func init() {
	flag.StringVar(&pathConfig, "config", confFile, "path config")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	flag.Parse()
	config.CreateConfig(pathConfig)
	conf := config.CreateConfig(pathConfig)
	logg := logger.CreateLogrusLog(conf)
	dbHandler := postgres.CreatePgConn(conf, logg)
	server(logg, dbHandler)
	scheduler(logg)

	<-stopCh

}

func server(log *logger.Logger, dbHandler *pg.DB) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer func() { _ = conn.Close() }()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer func() { _ = ch.Close() }()

	q, err := ch.QueueDeclare(
		"messages", // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := bodyFrom(os.Args)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Infof(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func scheduler(log *logger.Logger) {
	var errs []error

	crontab := cron.New()

	errs = append(errs, crontab.AddFunc("0 * * * * *", func() {
		tasks.ReportTask()
	}))

	crontab.Start()

	for _, err := range errs {
		if err != nil {
			log.Infof(err.Error())
		}
	}
	log.Info("Run scheduler")
}

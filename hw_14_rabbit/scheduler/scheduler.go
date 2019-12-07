package main

import (
	"flag"
	"github.com/robfig/cron"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/scheduler/config"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/scheduler/pkg/logger"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/scheduler/tasks"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
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
	flag.Parse()
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	config.CreateConfig(pathConfig)
	conf := config.CreateConfig(pathConfig)
	logg := logger.CreateLogrusLog(conf)
	conn := createRabbitConn(conf)
	defer func() { _ = conn.Close() }()
	ch := createChannel(conn)
	defer func() { _ = ch.Close() }()
	rk := "events"
	rabbitServer(logg, ch, rk)
	scheduler(logg, ch, rk)

	<-stopCh

}

func createRabbitConn(conf *config.Config) *amqp.Connection {
	strDial := "amqp://" + conf.RabbitUser + ":" + conf.RabbitPassword + "@localhost:" + conf.RabbitPort + "/"
	conn, err := amqp.Dial(strDial)
	failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func createChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch
}

func rabbitServer(log *logger.Logger, ch *amqp.Channel, rk string) {
	_, err := ch.QueueDeclare(
		rk,    // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")
}

func scheduler(log *logger.Logger, ch *amqp.Channel, rk string) {
	var errs []error
	crontab := cron.New()

	errs = append(errs, crontab.AddFunc("*/10 * * * * *", func() {
		tasks.EventReminder(log, ch, rk)
	}))

	crontab.Start()

	for _, err := range errs {
		if err != nil {
			log.Infof(err.Error())
		}
	}
	log.Info("Run scheduler")
}

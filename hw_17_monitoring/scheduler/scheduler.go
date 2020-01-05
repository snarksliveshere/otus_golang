package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/robfig/cron"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/scheduler/config"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/scheduler/pkg/logger"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/scheduler/tasks"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var conf config.AppConfig
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")
	logg := logger.CreateLogrusLog(conf.LogLevel)
	conn := createRabbitConn(&conf)
	stopCh := make(chan os.Signal, 1)
	defer func() { _ = conn.Close() }()
	ch := createChannel(conn)
	defer func() { _ = ch.Close() }()
	rk := "events"
	rabbitServer(logg, ch, rk)
	scheduler(logg, ch, rk, &conf)

	<-stopCh

}

func createRabbitConn(conf *config.AppConfig) *amqp.Connection {
	strDial := "amqp://" + conf.RbUser + ":" + conf.RbPassword + "@" + conf.RbHost + ":" + conf.RbPort + "/"
	for {
		conn, err := amqp.Dial(strDial)
		if err == nil {
			return conn
		} else {
			log.Printf("INFO:Failed to connect to RabbitMQ with %s", err.Error())
			time.Sleep(1 * time.Second)
		}
	}
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

func scheduler(log *logger.Logger, ch *amqp.Channel, rk string, conf *config.AppConfig) {
	var errs []error
	crontab := cron.New()

	errs = append(errs, crontab.AddFunc("*/10 * * * * *", func() {
		tasks.EventReminder(log, ch, rk, conf)
	}))

	crontab.Start()

	for _, err := range errs {
		if err != nil {
			log.Infof(err.Error())
		}
	}
	log.Info("Run scheduler")
}

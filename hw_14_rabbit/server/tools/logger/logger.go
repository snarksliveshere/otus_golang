package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/config"
)

const appName = "simple_app_calendar"

type Logger struct {
	log *logrus.Entry
}

func (logger *Logger) Infof(pattern string, args ...interface{}) {
	logger.log.Infof(pattern, args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.log.Info(args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.log.Fatal(args...)
}

func (logger *Logger) Fatalf(pattern string, args ...interface{}) {
	logger.log.Fatalf(pattern, args...)
}

func CreateLogrusLog(config *config.Config) *Logger {
	log := logrus.New()
	logEntry := logrus.NewEntry(log).WithField("app", appName)
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal("An error occurred during the logLevelAssertion")
	}
	log.SetLevel(level)

	return &Logger{log: logEntry}

	//return logEntry
}

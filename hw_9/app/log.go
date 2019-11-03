package app

import (
	"github.com/sirupsen/logrus"
	"sync"
)

const appName = "simple_app"

var (
	logOnce  sync.Once
	logEntry *logrus.Entry
)

func Log(path string) *logrus.Entry {
	logOnce.Do(func() {
		log := logrus.New()
		logEntry = logrus.NewEntry(log).WithField("app", appName)
		level, err := logrus.ParseLevel(Conf(path).LogLevel)
		if err != nil {
			log.Fatal("An error occurred during the logLevelAssertion")
		}
		log.SetLevel(level)
	})
	return logEntry
}

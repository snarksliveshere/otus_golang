package pkg

import (
	"github.com/sirupsen/logrus"
)

const appName = "simple_app_calendar"

type Logger struct {
	log logrus.Entry
}

func (logger *Logger) Log() logrus.Entry {
	return logger.log
}

func CreateLog(path string) Logger {

	log := logrus.New()
	logEntry := logrus.NewEntry(log).WithField("app", appName)
	level, err := logrus.ParseLevel(Conf(path).LogLevel)
	if err != nil {
		log.Fatal("An error occurred during the logLevelAssertion")
	}
	log.SetLevel(level)

	return Logger{log: *logEntry}

	//return logEntry
}

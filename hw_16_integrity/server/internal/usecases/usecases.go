package usecases

import (
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/server/entity"
)

type Logger interface {
	Info(args ...interface{})
	Infof(pattern string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(pattern string, args ...interface{})
}

type Actions struct {
	EventRepository entity.EventRepository
	DateRepository  entity.DateRepository
	Logger          Logger
}

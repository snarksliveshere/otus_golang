package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/entity"
)

type Logger interface {
	Info(args ...interface{})
	Infof(pattern string, args ...interface{})
}

type Actions struct {
	EventRepository entity.EventRepository
	DateRepository  entity.DateRepository
	Logger          Logger
}

package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_8/entity"
)

type Logger interface {
	Log(args ...interface{})
}

type Actions struct {
	EventRepository entity.EventRepository
	DateRepository  entity.DateRepository
	Logger          Logger
}

package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/entity"
)

type Logger interface {
	Info(args ...interface{})
	Infof(pattern string, args ...interface{})
}

type Actions struct {
	RecordRepository entity.RecordRepository
	DateRepository   entity.DateRepository
	Logger           Logger
}

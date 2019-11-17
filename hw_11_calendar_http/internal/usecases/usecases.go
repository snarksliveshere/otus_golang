package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/entity"
)

type Logger interface {
	Log(args ...interface{})
}

type Actions struct {
	RecordRepository entity.RecordRepository
	DateRepository   entity.DateRepository
	Logger           Logger
}

package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_8/src/entity"
)

type Logger interface {
	Log(args ...interface{})
}

type Actions struct {
	RecordRepository entity.RecordRepository
	DayRepository    entity.DayRepository
	Logger           Logger
}

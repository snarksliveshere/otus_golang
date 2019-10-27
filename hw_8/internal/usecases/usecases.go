package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_8/entity"
)

type Logger interface {
	Log(args ...interface{})
}

type Actions struct {
	RecordRepository entity.RecordRepository
	DateRepository   entity.DateRepository
	Logger           Logger
}

package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
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

package data_handlers

import (
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/internal/helpers"
	"time"
)

func CheckEventsForDay(date string) (time.Time, error) {
	return helpers.GetDateFromString(date)
}

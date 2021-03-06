package data_handlers

import (
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/internal/helpers"
	"time"
)

func CheckEventsForDay(date string) (time.Time, error) {
	return helpers.GetDateFromString(date)
}

package data_handlers

import (
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/helpers"
	"time"
)

func CheckEventsForDay(date string) (time.Time, error) {
	return helpers.GetDateFromString(date)
}

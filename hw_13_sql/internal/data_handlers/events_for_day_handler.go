package data_handlers

import (
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/helpers"
	"time"
)

func CheckEventsForDay(date string) (time.Time, error) {
	return helpers.GetDateFromString(date)
}

package data_handlers

import (
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/helpers"
	"time"
)

func CheckEventsForInterval(from, till string) (time.Time, time.Time, error) {
	tFrom, err := helpers.GetDateFromString(from)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	tTill, err := helpers.GetDateFromString(till)
	if err != nil {
		return tFrom, time.Time{}, err
	}

	return tFrom, tTill, nil
}

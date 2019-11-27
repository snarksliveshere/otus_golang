package data_handlers

import (
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/helpers"
	"time"
)

func CheckEventsForMonth(date string) (map[string]time.Time, error) {
	firstDate, err := helpers.GetFirstDateFromMonth(date)
	if err != nil {
		return nil, err
	}
	lastDate := firstDate.AddDate(0, 1, -1)
	m := make(map[string]time.Time, 2)
	m["firstDate"] = firstDate
	m["lastDate"] = lastDate

	return m, nil
}

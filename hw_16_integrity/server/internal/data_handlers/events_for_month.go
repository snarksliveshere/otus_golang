package data_handlers

import (
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/server/config"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/server/internal/helpers"
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

func CheckEventsForMonthString(date string) (map[string]string, error) {
	firstDate, err := helpers.GetFirstDateFromMonth(date)
	if err != nil {
		return nil, err
	}
	lastDate := firstDate.AddDate(0, 1, -1)
	m := make(map[string]string, 2)
	m["firstDate"] = firstDate.Format(config.TimeLayout)
	m["lastDate"] = lastDate.Format(config.TimeLayout)

	return m, nil
}

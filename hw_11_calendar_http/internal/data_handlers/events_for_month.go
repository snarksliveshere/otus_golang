package data_handlers

import (
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/helpers"
)

func CheckEventsForMonth(month string) (uint8, error) {
	return helpers.IsNumOfMonthInString(month)
}

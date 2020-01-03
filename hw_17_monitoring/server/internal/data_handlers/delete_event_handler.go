package data_handlers

import (
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/internal/helpers"
)

func CheckDeleteEvent(eventId string) (uint64, error) {
	n, err := validateDeleteEvent(eventId)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func validateDeleteEvent(eventId string) (uint64, error) {
	n, err := helpers.ParseStringToUint64(eventId)
	if err != nil {
		return 0, err
	}

	return n, nil
}

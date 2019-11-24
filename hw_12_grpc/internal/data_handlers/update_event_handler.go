package data_handlers

import (
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/internal/helpers"
)

func CheckUpdateEvent(title, desc string) (string, string, error) {
	err := validateUpdateEvent(title, desc)
	if err != nil {
		return title, desc, err
	}

	title, desc = modifierUpdateStringEvent(title, desc)
	return title, desc, nil
}

func validateUpdateEvent(title, desc string) error {
	if err := helpers.NotEmpty(title); err != nil {
		return err
	}
	if err := helpers.NotEmpty(desc); err != nil {
		return err
	}
	return nil
}

func modifierUpdateStringEvent(title, desc string) (string, string) {
	title = helpers.Trim(title)
	desc = helpers.Trim(desc)
	return title, desc
}

func ValidateUpdateEventId(eventId string) (uint64, error) {
	n, err := helpers.ParseStringToUint64(eventId)
	if err != nil {
		return 0, err
	}

	return n, nil
}

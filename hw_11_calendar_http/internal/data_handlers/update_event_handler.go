package data_handlers

import "github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/helpers"

func CheckUpdateEvent(title, desc, date, eventId string) (string, string, string, uint64, error) {
	err := validateUpdateEvent(title, desc, date, eventId)
	if err != nil {
		return title, desc, date, 0, err
	}

	return modifierUpdateEvent(title, desc, date, eventId)
}

func validateUpdateEvent(title, desc, date, eventId string) error {
	if err := helpers.NotEmpty(title); err != nil {
		return err
	}
	if err := helpers.NotEmpty(desc); err != nil {
		return err
	}
	if err := helpers.NotEmpty(date); err != nil {
		return err
	}
	if err := helpers.CanParseStringToUint64(eventId); err != nil {
		return err
	}
	return nil
}

func modifierUpdateEvent(title, desc, date, eventId string) (string, string, string, uint64, error) {
	title = helpers.Trim(title)
	desc = helpers.Trim(desc)
	date = helpers.Trim(date)
	n, _ := helpers.ParseStringToUint64(eventId)
	return title, desc, date, n, nil
}

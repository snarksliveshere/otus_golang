package data_handlers

import "github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/helpers"

func CheckCreateEvent(title, desc, date string) (string, string, string, error) {
	err := validateCreateEvent(title, desc, date)
	if err != nil {
		return title, desc, date, err
	}

	return modifierCreateEvent(title, desc, date)
}

func validateCreateEvent(title, desc, date string) error {
	if err := helpers.NotEmpty(title); err != nil {
		return err
	}
	if err := helpers.NotEmpty(desc); err != nil {
		return err
	}
	if err := helpers.NotEmpty(date); err != nil {
		return err
	}
	return nil
}

func modifierCreateEvent(title, desc, date string) (string, string, string, error) {
	title = helpers.Trim(title)
	desc = helpers.Trim(desc)
	date = helpers.Trim(date)
	return title, desc, date, nil
}
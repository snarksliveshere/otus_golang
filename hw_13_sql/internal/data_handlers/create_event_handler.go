package data_handlers

import (
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/helpers"
	"time"
)

func CheckCreateEvent(title, desc, date string) (string, string, time.Time, error) {
	err := validateCreateEvent(title, desc)
	if err != nil {
		return title, desc, time.Time{}, err
	}
	title, desc, err = modifierCreateEvent(title, desc)
	if err != nil {
		return title, desc, time.Time{}, err
	}
	day, err := helpers.GetDateFromString(date)
	if err != nil {
		return title, desc, time.Time{}, err
	}
	return title, desc, day, nil
}

func GetTimeFromString(date string) (time.Time, error) {
	return helpers.GetDateFromString(date)
}

func validateCreateEvent(title, desc string) error {
	if err := helpers.NotEmpty(title); err != nil {
		return err
	}
	if err := helpers.NotEmpty(desc); err != nil {
		return err
	}
	return nil
}

func modifierCreateEvent(title, desc string) (string, string, error) {
	title = helpers.Trim(title)
	desc = helpers.Trim(desc)
	return title, desc, nil
}

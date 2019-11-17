package validators

import "github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/helpers"

func CheckCreateEvent(title, desc string) (string, string, error) {
	err := validate(title, desc)
	if err != nil {
		return title, desc, err
	}

	return modifier(title, desc)
}

func validate(title, desc string) error {
	if err := helpers.NotEmpty(title); err != nil {
		return err
	}
	if err := helpers.NotEmpty(desc); err != nil {
		return err
	}
	return nil
}

func modifier(title, desc string) (string, string, error) {
	title = helpers.Trim(title)
	desc = helpers.Trim(desc)
	return title, desc, nil
}

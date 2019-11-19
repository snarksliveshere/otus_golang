package helpers

import (
	"errors"
	"strconv"
	"time"
)

func GetDateFromString(date string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, date)
	if err != nil {
		return t, err
	}
	return t, nil
}

func IsNumOfMonthInString(month string) (uint8, error) {
	i, err := strconv.ParseUint(month, 10, 64)
	if err != nil {
		return 0, err
	}
	if i >= 0 && i <= 12 {
		return uint8(i), nil
	}
	errs := errors.New("not valid month num")
	return 0, errs
}

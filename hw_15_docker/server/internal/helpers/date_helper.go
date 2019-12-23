package helpers

import (
	"errors"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/config"
	"strconv"
	"time"
)

func GetDateFromString(date string) (time.Time, error) {
	t, err := time.Parse(config.TimeLayout, date)
	if err != nil {
		return t, err
	}
	return t, nil
}

func GetDateTimeFromString(date string) (time.Time, error) {
	t, err := time.Parse(config.EventTimeLayout, date)
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

func MakeTimestampId() uint64 {
	return uint64(time.Now().UnixNano())
}

func GetFirstDateFromMonth(date string) (time.Time, error) {
	t, err := time.Parse(config.TimeMonthLayout, date)
	if err != nil {
		return t, err
	}
	return t, nil
}

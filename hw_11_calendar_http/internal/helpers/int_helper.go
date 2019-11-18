package helpers

import (
	"errors"
	"strconv"
)

func ParseToUint64(i int) (uint64, error) {
	var val float64
	err := errors.New("i can not parse this int to uint64")
	if val == float64(int(i)) {
		if i < 0 {
			return 0, err
		}
		return uint64(i), nil
	}
	return 0, err
}

func ParseStringToUint64(id string) (uint64, error) {
	i, err := strconv.ParseUint(id, 10, 64)
	if err == nil {
		return i, nil
	}

	return 0, err
}

func CanParseStringToUint64(id string) error {
	_, err := strconv.ParseUint(id, 10, 64)
	if err == nil {
		return nil
	}

	return err
}

package helpers

import "errors"

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

package helpers

import (
	"errors"
	"strings"
)

func NotEmpty(str string) (err error) {
	if str == "" {
		err = errors.New("empty string")
		return err
	}
	return nil
}

// только одна функция для примера. Ихх может быть очень много
func Trim(str string) string {
	return strings.TrimSpace(str)
}

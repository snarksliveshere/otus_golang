package unpackstr

import (
	"errors"
	"strconv"
	"strings"
)

const (
	errorsInvalidStr = "invalid incoming string structure"
)

func GetUnpackString(str string) (res string, err error) {
	if checkStrCorrect(str) == false {
		return res, errors.New(errorsInvalidStr)
	}
	storage := symbolDict(str)
	res = setUnpackSymbols(storage, str)

	return res, nil
}

func symbolDict(s string) (res []map[string]interface{}) {
	for i, v := range s {
		var prev string
		var slashes string
		vStr := string(v)
		m := make(map[string]interface{})
		if i != 0 {
			prev = string(s[i-1])
		}
		_, err := strconv.Atoi(vStr)

		if i > 1 && len(s) > 1 {
			slashes = s[i-2 : i]
		}
		if (vStr == `\` && prev != `\`) ||
			(err == nil && slashes == `\\`) ||
			(err == nil && prev != `\`) {
			continue
		}
		m["symbol"] = vStr
		m["index"] = i
		res = append(res, m)
	}
	return res
}

//func symbolFilter() bool {
//
//}

func setUnpackSymbols(storage []map[string]interface{}, s string) string {
	var str string
	for i, v := range storage {
		if i != (len(storage) - 1) {
			startVal := v["index"].(int)
			endVal := storage[i+1]["index"].(int)
			delta := endVal - startVal
			if delta >= 2 && s[startVal+1:endVal] != `\` {
				strNum := s[startVal+1 : endVal]
				strNum = strings.ReplaceAll(strNum, `\`, "")
				num, _ := strconv.Atoi(strNum)
				str += strings.Repeat(v["symbol"].(string), num)
			} else {
				str += v["symbol"].(string)
			}
			continue
		}
		val := v["index"].(int)
		if s[val+1:] != "" {
			num, _ := strconv.Atoi(s[val+1:])
			str += strings.Repeat(v["symbol"].(string), num)
		} else {
			str += v["symbol"].(string)
		}
	}

	return str
}

func checkStrCorrect(str string) bool {
	if str == "" {
		return false
	}
	_, errStr := strconv.Atoi(str)
	_, errFirstLetter := strconv.Atoi(string(str[0]))

	if errStr == nil || errFirstLetter == nil {
		return false
	}

	return true
}

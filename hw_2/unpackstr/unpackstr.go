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
	if !checkStrCorrect(str) {
		return res, errors.New(errorsInvalidStr)
	}
	storage := symbolDict(str)
	res = setUnpackSymbols(storage, str)

	return
}

func symbolDict(s string) (res []map[string]interface{}) {
	for i, v := range s {
		vStr := string(v)
		m := make(map[string]interface{})
		prev, slashes := getPrevSymbol(i, s)

		if skip := skipNotSymbol(vStr, prev, slashes); skip {
			continue
		}

		m["symbol"] = vStr
		m["index"] = i
		res = append(res, m)
	}
	return res
}

func getPrevSymbol(i int, s string) (prev, slashes string) {
	if i != 0 {
		prev = string(s[i-1])
	}
	if i > 1 && len(s) > 1 {
		slashes = s[i-2 : i]
	}

	return
}

func skipNotSymbol(el, prev, slashes string) bool {
	_, err := strconv.Atoi(el)
	if (el == `\` && prev != `\`) ||
		(err == nil && slashes == `\\`) ||
		(err == nil && prev != `\`) {
		return true
	}

	return false
}

func setUnpackSymbols(storage []map[string]interface{}, s string) (str string) {
	for i, v := range storage {
		if i != (len(storage) - 1) {
			startVal := v["index"].(int)
			endVal := storage[i+1]["index"].(int)
			str += getUnpackStrWithoutTail(v, s, startVal, endVal)
			continue
		}
		str += getTailUnpackStr(v, s)
	}

	return
}

func getUnpackStrWithoutTail(v map[string]interface{}, s string, startVal, endVal int) (str string) {
	delta := endVal - startVal
	if delta >= 2 && s[startVal+1:endVal] != `\` {
		strNum := s[startVal+1 : endVal]
		strNum = strings.ReplaceAll(strNum, `\`, "")
		num, err := strconv.Atoi(strNum)
		if err != nil {
			str = ""
		}
		str += strings.Repeat(v["symbol"].(string), num)
	} else {
		str += v["symbol"].(string)
	}
	return
}

func getTailUnpackStr(v map[string]interface{}, s string) (str string) {
	val := v["index"].(int)
	if s[val+1:] != "" {
		num, err := strconv.Atoi(s[val+1:])
		if err != nil {
			str = ""
		}
		str = strings.Repeat(v["symbol"].(string), num)
	} else {
		str = v["symbol"].(string)
	}
	return
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

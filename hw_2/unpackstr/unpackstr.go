package unpackstr

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetUnpackString(str string) (res string, err error) {
	if checkStrCorrect(str) == false {
		return res, errors.New("bad bad1")
	}

	storage := symbolDict(str)
	//fmt.Println(storage[0])

	setUnpackSymbols(storage, str)

	//numIndex := getAllNumIndex(str)
	//storage := createStorage(numIndex, str)
	//
	//if len(storage) == 0 {
	//	res := strings.ReplaceAll(str, `\`, "")
	//	return res, nil
	//}
	//
	//if res := returnUnpackStr(str, storage); res != "" {
	//	return res, nil
	//}
	//return res, errors.New("bad bad2")
	return "", nil
}

func symbolDict(s string) (res []map[string]interface{}) {
	for i, v := range s {
		var prev string
		m := make(map[string]interface{})
		if i != 0 {
			prev = string(s[i-1])
		}
		_, err := strconv.Atoi(string(v))
		if (string(v) == `\` && prev != `\`) ||
			(prev == `\` && err == nil && res[i-1]["symbol"] == `\`) ||
			(err == nil && prev != `\`) {
			continue
		}
		m["symbol"] = string(v)
		m["index"] = i
		res = append(res, m)
	}
	return res
}

func setUnpackSymbols(storage []map[string]interface{}, s string) {
	var str string
	fmt.Println(storage, str)
	for i, v := range storage {
		if i != (len(storage) - 1) {
			fmt.Println(len(storage), i)
			startVal := v["index"].(int)
			endVal := storage[i+1]["index"].(int)
			delta := endVal - startVal
			if delta >= 2 {
				strNum := s[startVal+1 : endVal]
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
	fmt.Println(str)
}

func returnUnpackStr(str string, storage map[int]string) string {
	var s string
	fmt.Println(str, storage)
	for i, v := range str {
		if _, err := strconv.Atoi(string(v)); err == nil {
			continue
		}
		if val, ok := storage[i]; ok {
			s += val
			fmt.Println(val, v)
			continue
		}
		s += string(v)
	}

	return s
}

func getAllNumIndex(str string) [][]int {
	re := regexp.MustCompile("[0-9]+")
	return re.FindAllStringIndex(str, -1)
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

func escapeSymbols(str string, v []int) (s string, a string, ind int, skip bool) {
	var prev string
	ind = v[0] - 1
	s = string(str[ind])
	a = str[v[0]:v[1]]
	if v[0] > 1 {
		prev = string(str[ind-1])
	}
	if s == `\` {
		if (v[1] - v[0]) <= 1 {
			if prev != `\` {
				skip = true
			}
			s = prev
		} else {
			a = str[v[0]+1 : v[1]]
			s = string(str[v[0]])
		}
	}
	return s, a, ind, skip
}

func createStorage(collection [][]int, str string) map[int]string {
	storage := make(map[int]string)
	for _, v := range collection {
		//ind := v[0] - 1
		multiSymbol, a, ind, skip := escapeSymbols(str, v)
		if skip == true {
			continue
		}
		fmt.Println(a)
		// if a0b2 == bb, then a01b2 == bb
		if string(a[0]) == "0" {
			storage[ind] = ""
			continue
		} else {
			num, er := strconv.Atoi(a)
			if er != nil {
				continue
			}
			storage[ind] = strings.Repeat(multiSymbol, num)
		}
	}
	fmt.Println(storage)
	return storage
}

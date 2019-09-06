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
	numIndex := getAllNumIndex(str)
	storage := createStorage(numIndex, str)

	if len(storage) == 0 {
		res := strings.ReplaceAll(str, "\\", "")
		return res, nil
	}

	if res := returnUnpackStr(str, storage); res != "" {
		return res, nil
	}
	return res, errors.New("bad bad2")
}

func returnUnpackStr(str string, storage map[int]string) string {
	var s string
	for i, v := range str {
		if _, err := strconv.Atoi(string(v)); err == nil {
			continue
		}
		if val, ok := storage[i]; ok {
			s += val
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

//func escapeSymbols(str string, v []int) (string, str) {
//	if (v[1] - v[0]) <= 1 {
//
//	}
//	a := str[v[0]+1:v[1]]
//	multiSymbol := string(str[v[0]])
//	return multiSymbol, a
//}

func createStorage(collection [][]int, str string) map[int]string {
	storage := make(map[int]string)
	for _, v := range collection {
		ind := v[0] - 1
		var a string
		multiSymbol := string(str[ind])
		fmt.Println(multiSymbol)
		if multiSymbol == "\\" {
			if (v[1] - v[0]) <= 1 {
				continue
			}
			a = str[v[0]+1 : v[1]]
			multiSymbol = string(str[v[0]])
		} else {
			a = str[v[0]:v[1]]
		}
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
	fmt.Println(collection)
	fmt.Println(storage)
	return storage
}

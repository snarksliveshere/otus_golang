package main

import (
	"fmt"
	"github.com/otus_golang/hw_2/unpackstr"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	storage := make(map[int]string)
	var s string

	if len(os.Args) != 2 {
		log.Fatal("olala")
	}
	str := os.Args[1]
	if unpackstr.CheckStrCorrect(str) == false {
		log.Fatal("oeee")
	}
	re := regexp.MustCompile("[0-9]+")
	collection := re.FindAllStringIndex(str, -1)

	for _, v := range collection {
		ind := v[0] - 1
		a := str[v[0]:v[1]]
		num, er := strconv.Atoi(a)
		if er != nil {
			continue
		}
		storage[ind] = strings.Repeat(string(str[ind]), num)
	}



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
	fmt.Println(s)

}

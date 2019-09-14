package main

import (
	"fmt"
	"github.com/otus_golang/hw_3/search"
	"io/ioutil"
	"log"
)

func main() {
	file, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	a := search.CommonWords(string(file), 10)
	fmt.Println(a)
}

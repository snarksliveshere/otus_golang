package main

import (
	"fmt"
	unpack "github.com/otus_golang/hw_2/unpackstr"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("olala")
	}
	str := os.Args[1]
	s, err := unpack.GetUnpackString(str)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(s)
}

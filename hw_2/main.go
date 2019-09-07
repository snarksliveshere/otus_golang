package main

import (
	"fmt"
	unpack "github.com/otus_golang/hw_2/unpackstr"
	"log"
	"os"
)

const (
	errorArgs = "there is no argument at the incoming point"
	success   = "program successfully completed"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal(errorArgs)
	}
	str := os.Args[1]
	s, err := unpack.GetUnpackString(str)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(s)
	log.Println(success)
}

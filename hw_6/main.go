package main

import (
	"flag"
	"fmt"
	"github.com/snarksliveshere/otus_golang/hw_6/gocopy"
	"log"
)

var (
	from, to      string
	offset, limit int64
	eof           string
)

func init() {
	flag.StringVar(&from, "from", "./files/if.txt", "if file")
	flag.StringVar(&to, "to", "./files/of.txt", "of file")
	flag.Int64Var(&offset, "offset", 0, "offset")
	flag.Int64Var(&limit, "limit", 0, "limit")
	flag.StringVar(&eof, "eof", "y", "eof [y/n]")
}

func main() {
	flag.Parse()
	err := gocopy.CopySubStr(from, to, limit, offset, eof)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Finish")
}

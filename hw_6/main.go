package main

import (
	"flag"
	"github.com/snarksliveshere/otus_golang/hw_6/gocopy"
)

var (
	from, to      string
	offset, limit int64
)

func init() {
	flag.StringVar(&from, "from", "./files/if.txt", "if file")
	flag.StringVar(&to, "to", "./files/of.txt", "of file")
	flag.Int64Var(&offset, "offset", 0, "offset")
	flag.Int64Var(&limit, "limit", 0, "limit")
}

func main() {
	flag.Parse()
	gocopy.CopySubStr(from, to, limit, offset)
}

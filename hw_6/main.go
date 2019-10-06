package main

import (
	"flag"
	"fmt"
	"os"
)

//const (
//	offset = 128
//	//ifPath = "./if.txt"
//	//ofPath = "./of.txt"
//	//limit  = 4
//)

var (
	from, to      string
	offset, limit int64
)

func init() {
	flag.StringVar(&from, "from", "./if.txt", "if file")
	flag.StringVar(&to, "to", "./of.txt", "of file")
	flag.Int64Var(&offset, "offset", 0, "offset")
	flag.Int64Var(&limit, "limit", 0, "limit")
}

func main() {
	flag.Parse()
	fmt.Println(from, to, offset, limit)
	ifFile, _ := os.Open(from)
	defer func() { _ = ifFile.Close() }()
	b := make([]byte, limit)
	ifFile.ReadAt(b, offset)
	fmt.Println(string(b))
	ofFile, _ := os.Create(to)
	ofFile.Write(b)

	//ifFile.Seek(offset, io.SeekStart)

	//read, _ := io.ReadFull(ifFile, b)
	//fmt.Println(string(b), read)
	defer func() { _ = ofFile.Close() }()

}

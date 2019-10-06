package main

import (
	"fmt"
	"os"
)

const (
	offset = 128
	ifPath = "./if.txt"
	ofPath = "./of.txt"
	limit  = 4
)

func main() {
	ifFile, _ := os.Open(ifPath)
	defer func() { _ = ifFile.Close() }()
	b := make([]byte, limit)
	ifFile.ReadAt(b, offset)
	fmt.Println(string(b))
	ofFile, _ := os.Create(ofPath)
	ofFile.Write(b)

	//ifFile.Seek(offset, io.SeekStart)

	//read, _ := io.ReadFull(ifFile, b)
	//fmt.Println(string(b), read)
	defer func() { _ = ofFile.Close() }()

}

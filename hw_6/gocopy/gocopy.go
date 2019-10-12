package gocopy

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
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

func CopySubStr() {
	flag.Parse()
	fmt.Println(from, to, offset, limit)
	//ifFile, err := os.Open(from)
	//fromPath, err := filepath.Abs(from)
	//errHandler(err)
	ifFile, err := os.Open(from)
	errHandler(err)
	ofFile, err := os.Create(to)
	errHandler(err)
	defer func() { _ = ifFile.Close() }()
	defer func() { _ = ofFile.Close() }()

	// reader section
	b := make([]byte, 0, limit)
	//_, err = ifFile.ReadAt(b, offset)
	pad := 10
	offs := offset
	for offset < (limit + offs) {
		fmt.Println(offset, limit)
		temp := make([]byte, pad)
		nBytes, err := ifFile.ReadAt(temp, offset)
		offset += int64(nBytes)
		if err == io.EOF {
			b = append(b, temp...)
			break
		}
		if err != nil {
			log.Panicf("failed to read: %v", err)
		}
		fmt.Printf("Length b %d  , temp %d\n", len(b), len(temp))

		b = append(b, temp...)
		fmt.Println(len(b), "lenB")
		fmt.Println(string(b))
	}
	//fmt.Println(len(b), "end")
	fs, _ := ifFile.Stat()
	fs.Size()
	fmt.Println(len(b), "end", fs.Size())

}

func errHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

package gocopy

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
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
	fromPath, err := filepath.Abs(from)
	errHandler(err)
	ifFile, err := os.Open(fromPath)
	errHandler(err)
	ofFile, err := os.Create(to)
	errHandler(err)

	defer func() { _ = ifFile.Close() }()
	defer func() { _ = ofFile.Close() }()

	// we will copy 200 Mb from /dev/rand to /dev/null
	//reader := io.LimitReader(rand.Reader, limit)
	//writer := ioutil.Discard
	//
	//// start new bar
	//bar := pb.Full.Start64(limit)
	//// create proxy reader
	//barReader := bar.NewProxyReader(reader)
	//// copy from proxy reader
	//io.Copy(writer, barReader)
	//// finish bar
	//bar.Finish()

	//runProgress(b, to)

	//time.Sleep(1 * time.Second)
	//b := make([]byte, limit)
	//_, err = ifFile.ReadAt(b, offset)
	//errHandler(err)
	//fmt.Println(string(b))
	//ch := make(chan int)
	//bw := bufio.NewWriter(ofFile)

	//runProgress(b)
	testIO(ifFile, ofFile, limit, offset)

	//pack := 12
	//var offs int
	//for offs < int(limit) {
	//	offs += pack
	//	if offs > len(b) {
	//		ch <- len(b) - (offs - pack)
	//		bw.Write(b[offs-pack:])
	//		break
	//	}
	//	ch <- pack
	//	bw.Write(b[offs-pack : offs])
	//
	//}
	//
	//bw.Flush()

	//
	//_, err = ofFile.Write(b)
	//errHandler(err)

	//ifFile.Seek(offset, io.SeekStart)

	//read, _ := io.ReadFull(ifFile, b)
	//fmt.Println(string(b))
}

func runProgress(b []byte) {
	//lenB := len(b)
	lenB := 100
	bar := pb.StartNew(lenB)
	for i := 0; i < lenB; i++ {
		//bar.Add(12)
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()

}

func testP() {
	count := 500
	// create and start new bar
	bar := pb.StartNew(count)

	// start bar from 'default' template
	// bar := pb.Default.Start(count)

	// start bar from 'simple' template
	// bar := pb.Simple.Start(count)

	// start bar from 'full' template
	// bar := pb.Full.Start(count)

	for i := 0; i < count; i++ {
		//bar.Increment()
		bar.Add(3)
		time.Sleep(time.Millisecond)
	}
	bar.Finish()
}

func testIO(ifFile *os.File, ofFile io.Writer, limit, offset int64) {
	b := make([]byte, limit)
	_, err := ifFile.ReadAt(b, offset)
	errHandler(err)
	reader := bytes.NewReader(b)
	// we will copy 200 Mb from /dev/rand to /dev/null
	//reader := io.LimitReader(rand.Reader, limit)
	//writer := ioutil.Discard

	// start new bar
	bar := pb.Full.Start64(limit)
	// create proxy reader
	//barReader := bar.NewProxyReader(io.Reader(&reader))
	barReader := bar.NewProxyReader(reader)
	// copy from proxy reader
	io.Copy(ofFile, barReader)
	// finish bar
	bar.Finish()
}
func errHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

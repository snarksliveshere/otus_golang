package gocopy

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var (
	clear map[string]func()
)

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		errHandler(err)
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func CopySubStr(from, to string, limit, offset int64) {
	ifFile, err := os.Open(from)
	errHandler(err)
	ofFile, err := os.Create(to)
	errHandler(err)
	defer func() { _ = ifFile.Close() }()
	defer func() { _ = ofFile.Close() }()

	pad := 10
	offs := offset
	bw := bufio.NewWriter(ofFile)
	for offset < (limit + offs) {
		if int(limit) < pad {
			pad = int(limit)
		}
		if (bw.Buffered() + pad) > int(limit) {
			pad = int(limit) - bw.Buffered()
		}
		if (bw.Buffered()+pad) > int(offset) && offset != 0 {
			pad = int(offset) - bw.Buffered()
		}
		temp := make([]byte, pad)
		nBytes, err := ifFile.ReadAt(temp, offset)
		offset += int64(nBytes)
		iter := offset - offs
		if err == io.EOF {
			CallClear()
			fmt.Printf("percent EOF %d  %%\n", iter*100/limit)
			if nBytes == 0 {
				errHandler(err)
			}
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("EOF file. Continue? [y/n]: ")
			text, errs := reader.ReadString('\n')
			errHandler(errs)
			if text == "y\n" {
				_, err = bw.Write(temp[:nBytes])
				errHandler(err)
				break
			}
			errHandler(err)
		}
		if err != nil {
			log.Panicf("failed to read: %v", err)
		}
		_, err = bw.Write(temp)
		errHandler(err)

		time.Sleep(10 * time.Millisecond)

		CallClear()
		fmt.Printf("percent %d  %%\n", iter*100/limit)

	}
	err = bw.Flush()
	errHandler(err)

	fs, _ := ofFile.Stat()
	fIf, _ := ifFile.Stat()

	fmt.Printf("Finish writing, dest file size: %d\n", fs.Size())
	fmt.Printf("src file size: %d\n", fIf.Size())
}

func errHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

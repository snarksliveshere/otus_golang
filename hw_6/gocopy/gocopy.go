package gocopy

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var (
	from, to      string
	offset, limit int64
	clear         map[string]func()
)

func init() {
	flag.StringVar(&from, "from", "./files/if.txt", "if file")
	flag.StringVar(&to, "to", "./files/of.txt", "of file")
	flag.Int64Var(&offset, "offset", 0, "offset")
	flag.Int64Var(&limit, "limit", 0, "limit")
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

func CopySubStr() {
	flag.Parse()
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
		fmt.Println(offset, bw.Buffered())
		if (bw.Buffered() + pad) > int(offset) {
			pad = int(offset) - bw.Buffered()
		}
		fmt.Println(pad, "pad")
		temp := make([]byte, pad)
		fmt.Println(offset, pad)
		nBytes, err := ifFile.ReadAt(temp, offset)

		if err == io.EOF {
			CallClear()
			fmt.Printf("percent EOF %d  %%\n", offset*100/limit)
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("EOF file. Continue? [y/n]: ")
			text, errs := reader.ReadString('\n')
			errHandler(errs)
			if text == "y\n" {
				_, err = bw.Write(temp)
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

		time.Sleep(100 * time.Millisecond)
		//CallClear()
		fmt.Printf("percent %d  %%\n", offset*100/limit)
		offset += int64(nBytes)

	}
	err = bw.Flush()
	errHandler(err)

	fmt.Println("Finish")
}

func errHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

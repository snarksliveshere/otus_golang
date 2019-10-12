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
		if err != nil {
			log.Fatalf(err.Error())
		}
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

func CopySubStr(from, to string, limit, offset int64, eof string) error {
	ifFile, err := os.Open(from)
	if err != nil {
		return err
	}
	ofFile, err := os.Create(to)
	if err != nil {
		return err
	}
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
				if err != nil {
					return err
				}
			}
			// TODO: пока не знаю, как в этом случае сделать лучше, но для тестов удобнее передать параметр
			//reader := bufio.NewReader(os.Stdin)
			//fmt.Print("EOF file. Continue? [y/n]: ")
			//text, errs := reader.ReadString('\n')
			//errHandler(errs)
			if eof == "y" {
				_, err = bw.Write(temp[:nBytes])
				if err != nil {
					return err
				}
				break
			}
			if err != nil {
				return err
			}
		}
		if err != nil {
			log.Panicf("failed to read: %v", err)
		}
		_, err = bw.Write(temp)
		if err != nil {
			return err
		}

		time.Sleep(20 * time.Millisecond)

		CallClear()
		fmt.Printf("percent %d  %%\n", iter*100/limit)

	}
	err = bw.Flush()
	if err != nil {
		return err
	}

	fOf, _ := ofFile.Stat()
	fIf, _ := ifFile.Stat()

	fmt.Printf("Finish writing, dest file size: %d\n", fOf.Size())
	fmt.Printf("src file size: %d\n", fIf.Size())
	return nil
}

package gocopy

import (
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
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func CopySubStr() {
	flag.Parse()
	//fmt.Println(from, to, offset, limit)
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
		//var percent int64
		if offset > (limit - int64(pad)) {
			pad = int(limit - offset)
		}
		temp := make([]byte, pad)
		nBytes, err := ifFile.ReadAt(temp, offset)
		offset += int64(nBytes)
		if err == io.EOF {
			b = append(b, temp...)
			CallClear()
			fmt.Printf("percent EOF %d  %%\n", offset*100/(limit*2))
			//reader := bufio.NewReader(os.Stdin)
			//fmt.Print("Enter text: ")
			//text, _ := reader.ReadString('\n')
			//fmt.Println(text)
			break
		}
		if err != nil {
			log.Panicf("failed to read: %v", err)
		}
		b = append(b, temp...)
		//ln := int64(len(b))
		//percent = offset * 100 / int64(ln)
		time.Sleep(10 * time.Millisecond)
		CallClear()

		fmt.Printf("percent %d  %%\n", offset*100/(limit*2))
	}
	//fmt.Println(len(b), "end")
	//fs, _ := ifFile.Stat()
	//fs.Size()
	//fmt.Println(len(b), "end", fs.Size())

}

func errHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

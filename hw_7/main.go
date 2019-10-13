package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("./prog").Output()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(out))

	//os.Setenv()
}

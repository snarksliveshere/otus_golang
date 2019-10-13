package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	files, err := ioutil.ReadDir("./envs")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name(), " f name")
		fo, err := os.Open("./envs/" + f.Name())
		if err != nil {
			log.Fatalf(err.Error())
		}
		r, err := ioutil.ReadAll(fo)
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("\nenv : %s,  var: %s\n", f.Name(), string(r))
		err = os.Setenv(f.Name(), string(r))
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
	fmt.Println("olla ", os.Getenv("first_env"))

	out, err := exec.Command("./test_prog/env_prog").Output()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(out))

	for _, f := range files {
		err := os.Unsetenv(f.Name())
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

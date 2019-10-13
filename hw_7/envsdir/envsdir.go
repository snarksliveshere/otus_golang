package envsdir

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func EnvsDir(pathEnvs, pathProg string) error {
	files, err := ioutil.ReadDir("./envs")
	if err != nil {
		return err
	}

	for _, f := range files {
		fmt.Println(f.Name(), " f name")
		fo, err := os.Open(pathEnvs + f.Name())
		if err != nil {
			return err
		}
		r, err := ioutil.ReadAll(fo)
		if err != nil {
			return err
		}
		fmt.Printf("\nenv : %s,  var: %s\n", f.Name(), string(r))
		err = os.Setenv(f.Name(), string(r))
		if err != nil {
			return err
		}
	}
	fmt.Println("olla ", os.Getenv("first_env"))

	out, err := exec.Command(pathProg).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))

	for _, f := range files {
		err := os.Unsetenv(f.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

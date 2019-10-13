package envsdir

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func EnvsDir(pathEnvs, pathProg string) error {
	files, err := ioutil.ReadDir(pathEnvs)
	if err != nil {
		return err
	}
	for _, f := range files {
		fo, err := os.Open(pathEnvs + f.Name())
		if err != nil {
			return err
		}
		r, err := ioutil.ReadAll(fo)
		if err != nil {
			return err
		}
		err = os.Setenv(f.Name(), string(r))
		if err != nil {
			return err
		}
	}
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

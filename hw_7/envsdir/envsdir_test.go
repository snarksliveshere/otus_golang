package envsdir

import (
	"os"
	"os/exec"
	"testing"
)

func TestExplicitError(t *testing.T) {
	cases := []struct {
		pEnv, pProg string
	}{
		{
			pEnv:  "../envs1/",
			pProg: "../test_prog/env_prog",
		},
		{
			pEnv:  "../envs/",
			pProg: "../test_prog/env_prog1",
		},
	}

	for _, c := range cases {
		err := EnvsDir(c.pEnv, c.pProg)
		if err == nil {
			t.Errorf("TestExplicitError(), here must be an error, env: %s, prog %s", c.pEnv, c.pProg)
		}
	}
}

func TestEnv(t *testing.T) {
	cases := []struct {
		pEnv, pProg, res string
	}{
		{
			pEnv:  "--path_prog=../envsdir/for_tests/for_tests",
			pProg: "--path_env_dir=../envs/",
			res:   "env_first_value",
		},
	}

	for _, c := range cases {
		args := []string{c.pEnv, c.pProg}
		out, err := exec.Command("../hw_7", args...).Output()
		if err != nil {
			t.Errorf("TestEnv(), err: %s, env: %s, prog %s", err.Error(), c.pEnv, c.pProg)
		}
		if c.res != string(out[0:len(out)-1]) {
			t.Errorf("TestEnv(), res: %s, out: %s", c.res, out)
		}
	}
}

func TestUnsetEnv(t *testing.T) {
	cases := []struct {
		pEnv, pProg, firstEnv, secondEnv string
	}{
		{
			pEnv:      "../envs/",
			pProg:     "../test_prog/env_prog",
			firstEnv:  "first_env",
			secondEnv: "second_env",
		},
	}

	for _, c := range cases {
		err := EnvsDir(c.pEnv, c.pProg)
		if err != nil {
			t.Errorf("TestUnsetEnv(), error: %s, env: %s, prog %s", err.Error(), c.pEnv, c.pProg)
		}
		if os.Getenv(c.firstEnv) != "" || os.Getenv(c.secondEnv) != "" {
			t.Errorf("TestUnsetEnv(), envs must be empty, first env: %s, second %s", os.Getenv(c.firstEnv), os.Getenv(c.secondEnv))
		}
	}
}

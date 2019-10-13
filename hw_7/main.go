package main

import (
	"flag"
	"github.com/snarksliveshere/otus_golang/hw_7/envsdir"
	"log"
)

var (
	pathEnvs, pathProg string
)

const (
	envDir     = "./envs/"
	pathToProg = "./test_prog/env_prog"
)

func init() {
	flag.StringVar(&pathEnvs, "path_env_dir", envDir, "path to env dir")
	flag.StringVar(&pathProg, "path_prog", pathToProg, "path to program")
}

func main() {
	flag.Parse()
	err := envsdir.EnvsDir(pathEnvs, pathProg)
	if err != nil {
		log.Fatal(err.Error())
	}
}

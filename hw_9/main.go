package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/snarksliveshere/otus_golang/hw_9/app"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	webServer  *http.Server
	pathConfig string
)

const (
	confFile = "./config/config.yaml"
)

func init() {
	flag.StringVar(&pathConfig, "config", confFile, "path config")
}

func main() {
	flag.Parse()
	stopch := make(chan os.Signal, 1)
	signal.Notify(stopch, syscall.SIGINT, syscall.SIGTERM)
	webApi()
	<-stopch
}

func webApi() {
	listenAddr := app.ListenAddr(app.Conf(pathConfig))
	router := mux.NewRouter()
	router.HandleFunc("/", helloHandler)
	webServer = &http.Server{Addr: listenAddr, Handler: router}

	go func() {
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Log(pathConfig).Panic(err)
		}
	}()

	app.Log(confFile).Infof("Run web webApi server: %v", listenAddr)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		app.Log(confFile).Fatal("An error occurred")
	}
}

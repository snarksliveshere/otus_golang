package main

import (
	"github.com/gorilla/mux"
	"github.com/snarksliveshere/otus_golang/hw_9/app"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	webServer *http.Server
)

func main() {
	stopch := make(chan os.Signal, 1)
	signal.Notify(stopch, syscall.SIGINT, syscall.SIGTERM)
	webApi()
	<-stopch
}

func webApi() {
	listenAddr := app.ListenAddr(app.Conf())
	router := mux.NewRouter()
	router.HandleFunc("/", helloHandler)
	webServer = &http.Server{Addr: listenAddr, Handler: router}

	go func() {
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Log().Panic(err)
		}
	}()

	app.Log().Infof("Run web webApi server: %v", listenAddr)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		app.Log().Fatal("An error occurred")
	}
}

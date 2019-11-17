package web

import (
	"github.com/gorilla/mux"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/cmd/inmem"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/config"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/pkg"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	log     *pkg.Logger
	storage *inmem.Storage
)

func Server(path string) {
	conf := config.CreateConfig(path)
	log = pkg.CreateLog(conf)

	storage = inmem.CreateStorageInstance(log)

	stopch := make(chan os.Signal, 1)
	signal.Notify(stopch, syscall.SIGINT, syscall.SIGTERM)
	webApi(conf)
	<-stopch
}

func webApi(conf *config.Config) {
	listenAddr := conf.ListenAddr()
	router := mux.NewRouter()
	routesRegister(router)
	webServer := &http.Server{Addr: listenAddr, Handler: router}
	go func() {
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	log.Infof("Run web webApi server: %v", listenAddr)
}

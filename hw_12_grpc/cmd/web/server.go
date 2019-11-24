package web

import (
	"github.com/gorilla/mux"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/cmd/inmem"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/config"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/pkg"
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

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	webApi(conf)
	<-stopCh
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

package web

import (
	"github.com/gorilla/mux"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/config"
	pg_repository2 "github.com/snarksliveshere/otus_golang/hw_15_docker/server/internal/interfaces/repositories/pg_repository"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/pkg/logger/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	log     *logrus.Logger
	storage *pg_repository2.Storage
)

func Server(path string) {
	conf := config.CreateConfig(path)
	log = logrus.CreateLogrusLog(conf)

	storage = pg_repository2.CreateStorageInstance(log, conf)

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

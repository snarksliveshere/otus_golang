package web

import (
	"github.com/gorilla/mux"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/server/config"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/server/internal/interfaces/repositories/pg_repository"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/server/internal/pkg/logger/logrus"
	"net/http"
)

var (
	log     *logrus.Logger
	storage *pg_repository.Storage
)

func Server(logg *logrus.Logger, conf *config.AppConfig) {
	log = logg
	storage = pg_repository.CreateStorageInstance(log, conf)
	webApi(conf)
}

func webApi(conf *config.AppConfig) {
	addr := conf.ListenIP + ":" + conf.WEBPort
	router := mux.NewRouter()
	routesRegister(router)

	webServer := &http.Server{Addr: addr, Handler: router}
	go func() {
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	log.Infof("Run web webApi server: %v", addr)
}

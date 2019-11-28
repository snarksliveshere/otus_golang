package web

import (
	"github.com/gorilla/mux"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/config"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/interfaces/repositories/mem_repository"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/tools/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	log     *logger.Logger
	storage *mem_repository.Storage
)

func Server(path string) {
	conf := config.CreateConfig(path)
	log = logger.CreateLogrusLog(conf)

	storage = mem_repository.CreateStorageInstance(log)

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

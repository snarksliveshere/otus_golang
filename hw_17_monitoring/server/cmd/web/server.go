package web

import (
	"github.com/alexellis/faas/gateway/metrics"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/config"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/internal/interfaces/repositories/pg_repository"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/internal/pkg/logger/logrus"
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
	metricsHandler := metrics.PrometheusHandler()
	router.Handle("/metrics", metricsHandler)
	webServer := &http.Server{Addr: addr, Handler: router}
	metricServer := &http.Server{Addr: ":9102", Handler: promhttp.Handler()}
	go func() {
		log.Info("metric server start")
		if err := metricServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}

	}()
	go func() {
		log.Info("web server start")
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}

	}()
	log.Infof("Run web webApi server: %v", addr)
}

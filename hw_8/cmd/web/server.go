package web

import (
	"github.com/gorilla/mux"
	"github.com/snarskliveshere/otus_golang/hw_8/config"
	"github.com/snarskliveshere/otus_golang/hw_8/pkg"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	log *pkg.Logger
)

func Server(path string) {
	conf := config.CreateConfig(path)
	log = pkg.CreateLog(conf)

	stopch := make(chan os.Signal, 1)
	signal.Notify(stopch, syscall.SIGINT, syscall.SIGTERM)
	webApi(conf)
	<-stopch
}

func webApi(conf *config.Config) {
	listenAddr := conf.ListenAddr()
	router := mux.NewRouter()
	router.HandleFunc("/", helloHandler)
	webServer := &http.Server{Addr: listenAddr, Handler: router}

	go func() {
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Log().Panic(err)
		}
	}()
	log.Log().Infof("Run web webApi server: %v", listenAddr)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Log().Fatal("An error occurred")
	}
}

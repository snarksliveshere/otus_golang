package web

import (
	"github.com/gorilla/mux"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/config"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/pkg"
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
	router.HandleFunc("/create-event", createEventHandler)
	router.HandleFunc("/update-event", updateEventHandler)
	router.HandleFunc("/delete-event", deleteEventHandler)
	router.HandleFunc("/events-for-day", eventsForDayHandler)
	router.HandleFunc("/events-for-week", eventsForWeekHandler)
	router.HandleFunc("/events-for-month", eventsForMonthHandler)

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

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Log().Fatal("An error occurred")
	}
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Log().Fatal("An error occurred")
	}
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Log().Fatal("An error occurred")
	}
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Log().Fatal("An error occurred")
	}
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Log().Fatal("An error occurred")
	}
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Log().Fatal("An error occurred")
	}
}

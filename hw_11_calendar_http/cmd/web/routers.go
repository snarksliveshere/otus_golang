package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

func routesRegister(router *mux.Router) {
	router.HandleFunc("/", helloHandler)
	router.HandleFunc("/create-event", createEventHandler)
	router.HandleFunc("/update-event", updateEventHandler)
	router.HandleFunc("/delete-event", deleteEventHandler)
	router.HandleFunc("/events-for-day", eventsForDayHandler)
	router.HandleFunc("/events-for-week", eventsForWeekHandler)
	router.HandleFunc("/events-for-month", eventsForMonthHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

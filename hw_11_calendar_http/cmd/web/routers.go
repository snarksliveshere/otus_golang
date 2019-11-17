package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/validators"
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

func notValidHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("400"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	a := -10
	b := uint64(a)
	fmt.Println(b)
	title, desc := r.FormValue("title"), r.FormValue("description")
	fmt.Println(title)
	title, desc, err := validators.CheckCreateEvent(title, desc)
	if err != nil {
		notValidHandler(w, r)
	}
	fmt.Println("olalal")
	fmt.Sprintf("%#v,%#v,%#v,", title, desc, err)
	//storage.AddRecord(title, desc)
	//
	//t := storage.FindRecordById(1)
	//
	//_, err = w.Write([]byte(t))
	//if err != nil {
	//	log.Fatal("An error occurred")
	//}
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

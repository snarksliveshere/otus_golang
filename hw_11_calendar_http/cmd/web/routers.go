package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/data_handlers"
	"net/http"
)

func routesRegister(router *mux.Router) {
	router.HandleFunc("/", helloHandler)
	router.HandleFunc("/create-event", validCreateEventHandler(createEventHandler))
	router.HandleFunc("/update-event", updateEventHandler)
	router.HandleFunc("/delete-event", deleteEventHandler)
	router.HandleFunc("/events-for-day", eventsForDayHandler)
	router.HandleFunc("/events-for-week", eventsForWeekHandler)
	router.HandleFunc("/events-for-month", eventsForMonthHandler)
}

func validCreateEventHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, desc, date := r.FormValue("title"), r.FormValue("description"), r.FormValue("date")
		title, desc, date, err := data_handlers.CheckCreateEvent(title, desc, date)
		if err != nil {
			notValidHandler(w, r)
			return
		}
		m := make(map[string]string, 3)
		m["title"] = title
		m["desc"] = desc
		m["date"] = date
		ctx := context.WithValue(r.Context(), "data", m)
		r = r.WithContext(ctx)
		h(w, r)
	}
}

func validDeleteEventHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventId := r.FormValue("eventId")
		n, err := data_handlers.CheckDeleteEvent(eventId)
		if err != nil {
			notValidHandler(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), "eventId", n)
		r = r.WithContext(ctx)
		h(w, r)
	}
}

func notValidHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("400"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func otherErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("500"))
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
	err := storage.AddRecord(title, desc, date)
	if err != nil {
		otherErrorHandler(w, r)
	}
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

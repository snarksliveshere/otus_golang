package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func routesRegister(router *mux.Router) {
	router.HandleFunc("/", helloHandler)
	router.HandleFunc("/create-event", validCreateEventHandler(createEventHandler)).Methods("POST")
	router.HandleFunc("/update-event", validUpdateEventHandler(updateEventHandler))
	router.HandleFunc("/delete-event", validDeleteEventHandler(deleteEventHandler))
	router.HandleFunc("/events-for-day", validEventsForDayHandler(eventsForDayHandler)).Queries("date", "{date}")
	router.HandleFunc("/events-for-week", eventsForWeekHandler)
	router.HandleFunc("/events-for-month", validEventsForMonthHandler(eventsForMonthHandler)).Queries("month", "{month}")
}

func notValidHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("400"))
	if err != nil {
		log.Fatal("not valid params", err.Error())
	}
}

func otherErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("500"))
	if err != nil {
		log.Fatal("An error occurred", err.Error())
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := ctx.Value("data").(map[string]string)
	date := ctx.Value("day").(time.Time)
	title, okTitle := data["title"]
	desc, okDesc := data["desc"]
	if !okTitle || !okDesc {
		otherErrorHandler(w, r)
	}
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
	ctx := r.Context()
	data := ctx.Value("data").(map[string]string)
	n := ctx.Value("eventId").(uint64)
	title, okTitle := data["title"]
	desc, okDesc := data["desc"]
	date, okDate := data["date"]
	fmt.Println(title, desc, date, n)

	if !okTitle || !okDesc || !okDate {
		otherErrorHandler(w, r)
	}
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	eventId := ctx.Value("data").(uint64)
	fmt.Println(eventId)

	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	date := ctx.Value("date").(time.Time)
	fmt.Println(date)
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	month := ctx.Value("month").(uint8)
	fmt.Println(month)
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

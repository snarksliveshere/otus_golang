package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/entity"
	"net/http"
	"time"
)

const (
	statusOK    = "success"
	statusError = "error"
)

type Response struct {
	Date       entity.Date   `json:"day,omitempty"`
	Record     entity.Record `json:"record,omitempty"`
	Collection []interface{} `json:"collection,omitempty"`
	//Result     []string      `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
	Status string `json:"status,omitempty"`
}

//curl -d 'title=some-title&description=some_desc&date=2019-11-01' -X POST http://localhost:3001/create-event

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
	date := ctx.Value("date").(time.Time)
	title, okTitle := data["title"]
	desc, okDesc := data["desc"]
	if !okTitle || !okDesc {
		otherErrorHandler(w, r)
	}
	rec, day, err := storage.AddRecord(title, desc, date)
	resp := Response{Date: day, Record: rec, Status: statusOK}
	jResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jResp)
	if err != nil {
		otherErrorHandler(w, r)
	}
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
	eventId := ctx.Value("eventId").(uint64)
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

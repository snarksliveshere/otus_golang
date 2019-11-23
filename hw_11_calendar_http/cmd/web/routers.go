package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/config"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/entity"
	"net/http"
	"time"
)

const (
	statusOK    = "success"
	statusError = "error"
)

type Response struct {
	Date       entity.Date     `json:"day,omitempty"`
	Record     entity.Record   `json:"record,omitempty"`
	Records    []entity.Record `json:"records,omitempty"`
	Collection []interface{}   `json:"collection,omitempty"`
	//Result     []string      `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
	Status string `json:"status,omitempty"`
}

func routesRegister(router *mux.Router) {
	router.HandleFunc("/", helloHandler)
	router.HandleFunc("/create-event", validCreateEventHandler(createEventHandler)).Methods("POST")
	router.HandleFunc("/update-event", validUpdateEventHandler(updateEventHandler))
	router.HandleFunc("/delete-event", validDeleteEventHandler(deleteEventHandler))
	router.HandleFunc("/events-for-day", validEventsForDayHandler(eventsForDayHandler)).Queries("date", "{date}")
	router.HandleFunc("/events-for-week", eventsForWeekHandler)
	router.HandleFunc("/events-for-month", validEventsForMonthHandler(eventsForMonthHandler)).Queries("month", "{month}")
	//router.HandleFunc("/events-for-month", eventsForMonthHandler).Queries("month", "{month}")
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

//curl -d 'title=some-title&description=some_desc&date=2019-11-01' -X POST http://localhost:3001/create-event
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := ctx.Value("data").(map[string]string)
	date, okDate := ctx.Value("date").(time.Time)
	title, okTitle := data["title"]
	desc, okDesc := data["desc"]
	if !okTitle || !okDesc || !okDate {
		otherErrorHandler(w, r)
	}
	rec, day, c, err := storage.AddRecord(title, desc, date)
	fmt.Println(c)
	if err != nil {
		otherErrorHandler(w, r)
	}
	resp := Response{Date: *day, Record: rec, Status: statusOK}
	sendResponse(resp, w, r)
}

// curl -d 'title=new-title&description=new_desc&date=2019-11-01&eventId=123' -X POST http://localhost:3001/update-event
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := ctx.Value("data").(map[string]string)
	n, okN := ctx.Value("eventId").(uint64)
	date, okDate := ctx.Value("date").(time.Time)
	title, okTitle := data["title"]
	desc, okDesc := data["desc"]

	if !okTitle || !okDesc || !okDate || !okN {
		otherErrorHandler(w, r)
	}
	err := storage.UpdateRecordById(n, date, title, desc)
	resp := Response{}
	if err != nil {
		resp.Status = statusError
		resp.Error = err.Error()
		otherErrorHandler(w, r)
	} else {
		resp.Status = statusOK
	}
	sendResponse(resp, w, r)
	_, err = w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

//curl -d 'eventId=123' -X POST http://localhost:3001/delete-event
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	eventId, ok := ctx.Value("eventId").(uint64)
	if !ok {
		otherErrorHandler(w, r)
	}
	err := storage.DeleteRecordById(eventId)
	resp := Response{}
	if err != nil {
		resp.Status = statusError
		resp.Error = err.Error()
		otherErrorHandler(w, r)
	} else {
		resp.Status = statusOK
	}
	sendResponse(resp, w, r)
}

func sendResponse(resp Response, w http.ResponseWriter, r *http.Request) {
	jResp, err := json.Marshal(resp)
	if err != nil {
		otherErrorHandler(w, r)
	}
	w.Header().Set(config.HeaderContentType, config.HeaderContentType)
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jResp)
	if err != nil {
		otherErrorHandler(w, r)
	}
}

// curl 'http://localhost:3001/events-for-day?date=2019-11-01'
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	date, okDate := ctx.Value("date").(time.Time)

	if !okDate {
		otherErrorHandler(w, r)
	}
	day, err := storage.GetEventsForDay(date)
	resp := Response{}
	if err != nil {
		resp.Status = statusError
		resp.Error = err.Error()
		otherErrorHandler(w, r)
	} else {
		resp.Status = statusOK
		resp.Records = day.Records
	}
	sendResponse(resp, w, r)
}

// curl 'http://localhost:3001/events-for-month?month=2019-11'
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dates, okDates := ctx.Value("dates").(map[string]time.Time)
	if !okDates {
		otherErrorHandler(w, r)
	}
	fmt.Println(dates)
	t, _ := time.Parse("2006-01", "2019-11")
	fmt.Println(t)
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

// curl 'http://localhost:3001/events-for-week?date=2019-11-01'
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

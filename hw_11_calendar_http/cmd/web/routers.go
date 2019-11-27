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
	Date       entity.Date    `json:"day,omitempty"`
	Event      entity.Event   `json:"event,omitempty"`
	Events     []entity.Event `json:"events,omitempty"`
	Collection []interface{}  `json:"collection,omitempty"`
	//Result     []string      `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
	Status string `json:"status,omitempty"`
}

func routesRegister(router *mux.Router) {
	router.HandleFunc("/", helloHandler)
	router.HandleFunc("/create-event", validCreateEventHandler(createEventHandler)).Methods("POST")
	router.HandleFunc("/update-event", validUpdateEventHandler(updateEventHandler))
	router.HandleFunc("/delete-event", validDeleteEventHandler(deleteEventHandler))
	router.HandleFunc("/events-for-day", validEventsForDayHandler(eventsForDayHandler)).
		Queries("date", "{date}")
	router.HandleFunc("/events-for-week", validEventsForIntervalHandler(eventsForWeekHandler)).
		Queries("from", "{from}", "till", "{till}")
	router.HandleFunc("/events-for-month", validEventsForMonthHandler(eventsForMonthHandler)).
		Queries("month", "{month}")
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
	rec, day, c, err := storage.AddEvent(title, desc, date)
	fmt.Println(c)
	if err != nil {
		otherErrorHandler(w, r)
	}
	resp := Response{Date: *day, Event: rec, Status: statusOK}
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
	err := storage.UpdateEventById(n, date, title, desc)
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
	err := storage.DeleteEventById(eventId)
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

//curl -d 'title=some-title&description=some_desc&date=2019-11-01' -X POST http://localhost:3001/create-event
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
		resp.Events = day.Events
	}
	sendResponse(resp, w, r)
}

//curl -d 'title=some-title&description=some_desc&date=2019-11-01' -X POST http://localhost:3001/create-event
// curl 'http://localhost:3001/events-for-month?month=2019-11'
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dates, okDates := ctx.Value("dates").(map[string]time.Time)
	if !okDates {
		otherErrorHandler(w, r)
	}
	events, err := storage.GetEventsForInterval(dates["firstDate"], dates["lastDate"])
	resp := Response{}
	if err != nil {
		resp.Status = statusError
		resp.Error = err.Error()
		otherErrorHandler(w, r)
	} else {
		resp.Status = statusOK
		resp.Events = events
	}
	sendResponse(resp, w, r)
}

// curl -d 'title=some-title&description=some_desc&date=2019-11-01' -X POST http://localhost:3001/create-event
// curl 'http://localhost:3001/events-for-week?from=2019-11-01&till=2019-11-08'
// кто там неделя, это решает передающий данные, я получаю их как интервал
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dates, okDates := ctx.Value("dates").(map[string]time.Time)
	if !okDates {
		otherErrorHandler(w, r)
	}
	events, err := storage.GetEventsForInterval(dates["from"], dates["till"])
	resp := Response{}
	if err != nil {
		resp.Status = statusError
		resp.Error = err.Error()
		otherErrorHandler(w, r)
	} else {
		resp.Status = statusOK
		resp.Events = events
	}
	sendResponse(resp, w, r)
}

package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/config"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/data_handlers"
	"net/http"
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
	router.HandleFunc("/create-event", validCreateEventHandler(createEventHandler)).Methods(http.MethodPost)
	router.HandleFunc("/update-event", validUpdateEventHandler(updateEventHandler))
	router.HandleFunc("/delete-event", validDeleteEventHandler(deleteEventHandler))
	router.HandleFunc("/events-for-day", validEventsForDayHandler(eventsForDayHandler)).
		Queries("date", "{date}")
	router.HandleFunc("/events-for-week", validEventsForIntervalHandler(eventsForWeekHandler)).
		Queries("from", "{from}", "till", "{till}")
	router.HandleFunc("/events-for-month", validEventsForMonthHandler(eventsForMonthHandler)).
		Queries("month", "{month}")
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
	title, desc, date := r.FormValue("title"), r.FormValue("description"), r.FormValue("date")
	title, desc, day, err := data_handlers.CheckCreateEvent(title, desc, date)
	if err != nil {
		notValidHandler(w, r)
		return
	}
	rec, d, c, err := storage.AddRecord(title, desc, day)
	fmt.Println(c)
	if err != nil {
		otherErrorHandler(w, r)
	}
	resp := Response{Date: *d, Record: rec, Status: statusOK}
	sendResponse(resp, w, r)
}

// curl -d 'title=new-title&description=new_desc&date=2019-11-01&eventId=123' -X POST http://localhost:3001/update-event
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	title, desc, date, eventId := r.FormValue("title"),
		r.FormValue("description"),
		r.FormValue("date"),
		r.FormValue("eventId")
	title, desc, day, id, err := data_handlers.CheckUpdateEvent(title, desc, date, eventId)
	if err != nil {
		notValidHandler(w, r)
		return
	}

	err = storage.UpdateRecordById(id, day, title, desc)
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

//curl -d 'eventId=123' -X POST http://localhost:3001/delete-event
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	eventId := r.FormValue("eventId")
	id, err := data_handlers.CheckDeleteEvent(eventId)
	if err != nil {
		notValidHandler(w, r)
		return
	}
	err = storage.DeleteRecordById(id)
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

//curl -d 'title=some-title&description=some_desc&date=2019-11-01' -X POST http://localhost:3001/create-event
// curl 'http://localhost:3001/events-for-day?date=2019-11-01'
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date, ok := vars["date"]
	if !ok {
		notValidHandler(w, r)
		return
	}
	t, err := data_handlers.CheckEventsForDay(date)
	if err != nil {
		notValidHandler(w, r)
		return
	}
	day, err := storage.GetEventsForDay(t)

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

//curl -d 'title=some-title&description=some_desc&date=2019-11-01' -X POST http://localhost:3001/create-event
// curl 'http://localhost:3001/events-for-month?month=2019-11'
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	month, ok := vars["month"]
	if !ok {
		notValidHandler(w, r)
		return
	}
	dates, err := data_handlers.CheckEventsForMonth(month)
	if err != nil {
		notValidHandler(w, r)
		return
	}

	records, err := storage.GetEventsForInterval(dates["firstDate"], dates["lastDate"])
	resp := Response{}
	if err != nil {
		resp.Status = statusError
		resp.Error = err.Error()
		otherErrorHandler(w, r)
	} else {
		resp.Status = statusOK
		resp.Records = records
	}
	sendResponse(resp, w, r)
}

// curl -d 'title=some-title&description=some_desc&date=2019-11-01' -X POST http://localhost:3001/create-event
// curl 'http://localhost:3001/events-for-week?from=2019-11-01&till=2019-11-08'
// кто там неделя, это решает передающий данные, я получаю их как интервал
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	from, okFrom := vars["from"]
	till, okTill := vars["till"]
	if !okFrom || !okTill {
		notValidHandler(w, r)
		return
	}
	tFrom, tTill, err := data_handlers.CheckEventsForInterval(from, till)

	records, err := storage.GetEventsForInterval(tFrom, tTill)
	resp := Response{}
	if err != nil {
		resp.Status = statusError
		resp.Error = err.Error()
		otherErrorHandler(w, r)
	} else {
		resp.Status = statusOK
		resp.Records = records
	}
	sendResponse(resp, w, r)
}

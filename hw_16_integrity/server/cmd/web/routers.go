package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/server/config"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/server/entity"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/server/internal/data_handlers"
	"net/http"
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
	router.HandleFunc("/healthcheck", healthCheckHandler)
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

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		log.Fatal("An error occurred")
	}
}

//curl -d 'title=some-title&description=some_desc&date=2019-11-01' -X POST http://localhost:3001/create-event
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	title, desc, date := r.FormValue("title"), r.FormValue("description"), r.FormValue("date")
	title, desc, time, err := data_handlers.CheckCreateEvent(title, desc, date)
	resp := Response{}
	if err != nil {
		resp.Status = statusError
		resp.Error = err.Error()
	}
	recId, err := storage.Actions.CreateEvent(title, desc, time)
	if err != nil {
		resp.Status = statusError
		resp.Error = err.Error()
	}
	rec, err := storage.Actions.EventRepository.FindById(recId)
	if err != nil {
		resp.Status = statusError
		resp.Error = err.Error()
	} else {
		resp.Status = statusOK
		resp.Event = rec
	}
	sendResponse(resp, w, r)
}

// curl -d 'title=new-title&description=new_desc&date=2019-11-01&eventId=123' -X POST http://localhost:3001/update-event
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	title, desc, date, eventId := r.FormValue("title"),
		r.FormValue("description"),
		r.FormValue("date"),
		r.FormValue("eventId")
	title, desc, _, id, err := data_handlers.CheckUpdateEvent(title, desc, date, eventId)
	if err != nil {
		notValidHandler(w, r)
		return
	}

	err = storage.Actions.UpdateEventById(id, title, desc)
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
	err = storage.Actions.DeleteEventById(id)
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

	events, err := storage.Actions.GetEventsByDay(date)

	resp := Response{}
	if err != nil {
		resp.Status = statusError
		resp.Error = err.Error()
	} else {
		resp.Status = statusOK
		resp.Events = events
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
	dates, err := data_handlers.CheckEventsForMonthString(month)
	if err != nil {
		notValidHandler(w, r)
		return
	}

	events, err := storage.Actions.EventRepository.GetEventsByDateInterval(dates["firstDate"], dates["lastDate"])
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
	vars := mux.Vars(r)
	from, okFrom := vars["from"]
	till, okTill := vars["till"]
	if !okFrom || !okTill {
		notValidHandler(w, r)
		return
	}
	//tFrom, tTill, err := data_handlers.CheckEventsForInterval(from, till)

	events, err := storage.Actions.EventRepository.GetEventsByDateInterval(from, till)
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

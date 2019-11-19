package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/data_handlers"
	"net/http"
)

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
		log.Infof("create-event query with %#v", m)
		h(w, r)
	}
}

func validUpdateEventHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, desc, date, eventId := r.FormValue("title"),
			r.FormValue("description"),
			r.FormValue("date"),
			r.FormValue("eventId")

		title, desc, date, n, err := data_handlers.CheckUpdateEvent(title, desc, date, eventId)
		if err != nil {
			notValidHandler(w, r)
			return
		}
		m := make(map[string]string, 4)
		m["title"] = title
		m["desc"] = desc
		m["date"] = date
		ctx := context.WithValue(r.Context(), "data", m)
		ct := context.WithValue(ctx, "eventId", n)
		r = r.WithContext(ct)
		log.Infof("update-event query with %#v, %v", m, n)
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
		log.Infof("delete-event query with %v", n)
		h(w, r)
	}
}

func validEventsForDayHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		date, ok := vars["month"]
		if !ok {
			notValidHandler(w, r)
			return
		}
		t, err := data_handlers.CheckEventsForDay(date)
		if err != nil {
			notValidHandler(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), "date", t)
		r = r.WithContext(ctx)
		log.Infof("events-for-day query with date %v", t)
		h(w, r)
	}
}

func validEventsForMonthHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		month, ok := vars["month"]
		if !ok {
			notValidHandler(w, r)
			return
		}
		n, err := data_handlers.CheckEventsForMonth(month)
		if err != nil {
			notValidHandler(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), "month", n)
		r = r.WithContext(ctx)
		log.Infof("events-for-month query with date %v", n)
		h(w, r)
	}
}
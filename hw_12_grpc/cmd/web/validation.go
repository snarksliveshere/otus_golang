package web

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/internal/data_handlers"
	"net/http"
	"time"
)

func validCreateEventHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, desc, date := r.FormValue("title"), r.FormValue("description"), r.FormValue("date")
		title, desc, err := data_handlers.CheckCreateEvent(title, desc)
		if err != nil {
			notValidHandler(w, r)
			return
		}
		day, err := data_handlers.GetTimeFromString(date)
		if err != nil {
			notValidHandler(w, r)
			return
		}
		m := make(map[string]string, 3)
		m["title"] = title
		m["desc"] = desc
		ctx := context.WithValue(r.Context(), "data", m)
		ct := context.WithValue(ctx, "date", day)
		r = r.WithContext(ct)
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

		title, desc, err := data_handlers.CheckUpdateEvent(title, desc)
		if err != nil {
			notValidHandler(w, r)
			return
		}
		n, err := data_handlers.ValidateUpdateEventId(eventId)
		if err != nil {
			notValidHandler(w, r)
			return
		}
		day, err := data_handlers.GetTimeFromString(date)
		if err != nil {
			notValidHandler(w, r)
			return
		}

		m := make(map[string]string, 4)
		m["title"] = title
		m["desc"] = desc
		ctx := context.WithValue(r.Context(), "data", m)
		ct := context.WithValue(ctx, "eventId", n)
		c := context.WithValue(ct, "date", day)
		r = r.WithContext(c)
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
		m, err := data_handlers.CheckEventsForMonth(month)
		if err != nil {
			notValidHandler(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), "dates", m)
		r = r.WithContext(ctx)
		log.Infof("events-for-month query with date %v", m)
		h(w, r)
	}
}

func validEventsForIntervalHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		from, okFrom := vars["from"]
		till, okTill := vars["till"]
		if !okFrom || !okTill {
			notValidHandler(w, r)
			return
		}
		tFrom, err := data_handlers.CheckEventsForDay(from)
		tTill, err := data_handlers.CheckEventsForDay(till)
		if err != nil {
			notValidHandler(w, r)
			return
		}
		mt := make(map[string]time.Time, 2)
		mt["from"] = tFrom
		mt["till"] = tTill
		ctx := context.WithValue(r.Context(), "dates", mt)
		r = r.WithContext(ctx)
		log.Infof("events-for-day query with date %v", mt)
		h(w, r)
	}
}

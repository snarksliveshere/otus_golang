package web

import (
	"context"
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

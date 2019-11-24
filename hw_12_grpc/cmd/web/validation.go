package web

import (
	"github.com/gorilla/mux"
	"net/http"
)

func validCreateEventHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, desc, date := r.FormValue("title"), r.FormValue("description"), r.FormValue("date")
		log.Infof("create-event query with title: %#v, desc: %#v, date: %#v", title, desc, date)
		h(w, r)
	}
}

func validUpdateEventHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, desc, date, eventId := r.FormValue("title"),
			r.FormValue("description"),
			r.FormValue("date"),
			r.FormValue("eventId")
		log.Infof("update-event query with title: %#v, desc: %v, date: %#v, eventId: %#v", title, desc, date, eventId)
		h(w, r)
	}
}

func validDeleteEventHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventId := r.FormValue("eventId")
		log.Infof("delete-event query with %v", eventId)
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
		log.Infof("events-for-day query with date %v", date)
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
		log.Infof("events-for-month query with date %v", month)
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
		log.Infof("events-for-interval query with from date: %v and till date: %#v", from, till)
		h(w, r)
	}
}

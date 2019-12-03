package entity

import "time"

type Calendar struct {
	Date
}

type Date struct {
	Day    time.Time `json:"day,omitempty"`
	Events []Event   `json:"events,omitempty"`
}

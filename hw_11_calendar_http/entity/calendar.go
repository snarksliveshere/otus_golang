package entity

import "time"

type Calendar struct {
	Dates []Date `json:"dates,omitempty"`
}

type Date struct {
	Day     time.Time `json:"day,omitempty"`
	Records []Record  `json:"records,omitempty"`
}

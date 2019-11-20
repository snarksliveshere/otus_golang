package entity

import "time"

type Calendar struct {
	Dates []Date
}

type Date struct {
	Day     time.Time
	Records []Record
}

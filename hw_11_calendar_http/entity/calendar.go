package entity

import "time"

type Calendar struct {
	Date
}

type Date struct {
	Day     time.Time
	Records []Record
}

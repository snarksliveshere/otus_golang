package entity

import "time"

type Calendar struct {
	Date
}

type Date struct {
	Id      uint64
	Item    time.Time
	Records []Record
}

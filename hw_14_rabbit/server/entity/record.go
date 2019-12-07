package entity

import "time"

type Event struct {
	Id          uint64    `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Time        time.Time `json:"time"`
	DateFk      uint32    `json:"dateFk"`
}

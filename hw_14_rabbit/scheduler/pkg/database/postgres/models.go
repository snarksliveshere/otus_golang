package postgres

import "time"

type Event struct {
	TableName   struct{} `sql:"calendar.event"`
	Id          uint64
	Title       string    `sql:"title, notnull"`
	Description string    `sql:"description"`
	Time        time.Time `sql:"time,notnull,unique:public_event_time_date_uidx"`
	DateFk      uint32    `sql:"date_fk,notnull,unique:public_event_time_date_uidx"`
}

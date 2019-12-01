package entity

import (
	"github.com/go-pg/pg"
	"time"
)

type RecordRepository interface {
	FindById(id uint64) (Record, error)
	GetEventsByDay(dayFk uint32) ([]Record, error)
	Save(record Record) (pg.Result, error)
	//Save(record Record) (uint64, error)

	Delete(record Record) error
	Edit(record Record) error
	Show() []Record
}

type DateRepository interface {
	FindByDay(day string) (Date, error)
	Save(date Date) (uint32, error)

	AddRecordToDate(record Record, day *Date) error
	AddDateToCalendar(day Date) error
	ShowDayRecords(day *Date) ([]Record, error)
	GetDateFromString(date string) (time.Time, error)
	GetCalendar() *Calendar
}

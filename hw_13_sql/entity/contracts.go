package entity

import "time"

type RecordRepository interface {
	FindById(id uint64) (Record, error)
	GetEventsByDay(dayFk uint32) ([]Record, error)

	Delete(record Record) error
	Edit(record Record) error
	Show() []Record
	Save(record Record) error
}

type DateRepository interface {
	FindByDay(day string) (Date, error)

	AddRecordToDate(record Record, day *Date) error
	AddDateToCalendar(day Date) error
	ShowDayRecords(day *Date) ([]Record, error)
	GetDateFromString(date string) (time.Time, error)
	GetCalendar() *Calendar
}

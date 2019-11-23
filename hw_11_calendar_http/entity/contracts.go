package entity

import "time"

type RecordRepository interface {
	FindById(id uint64) (Record, error)
	Delete(record Record) error
	Edit(record Record) error
	Show() []Record
	Save(record Record) error
}

type DateRepository interface {
	AddRecordToDate(record Record, day *Date) error
	AddDateToCalendar(day Date) error
	ShowDayRecords(day *Date) ([]Record, error)
	FindByDay(day time.Time, c *Calendar) (*Date, error)
	GetDateFromString(date string) (time.Time, error)
	GetCalendar() *Calendar
}

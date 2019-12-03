package entity

import "time"

type EventRepository interface {
	FindById(id uint64) (Event, error)
	Delete(event Event) error
	Edit(event Event) error
	Show() []Event
	Save(event Event) error
}

type DateRepository interface {
	AddEventToDate(event Event, day *Date) error
	AddDateToCalendar(day Date) error
	ShowDayEvents(day *Date) ([]Event, error)
	FindByDay(day time.Time, c *Calendar) (*Date, error)
	GetDateFromString(date string) (time.Time, error)
	GetCalendar() *Calendar
}

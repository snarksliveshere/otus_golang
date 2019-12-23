package entity

import "time"

type EventRepository interface {
	FindById(id uint64) (Event, error)
	GetEventsByDay(dayFk uint32) ([]Event, error)
	Save(event Event) (uint64, error)
	Delete(event Event) error
	Edit(event Event) error
	GetEventsByDateInterval(from, till string) ([]Event, error)
	GetEventsByTimeInterval(from, till time.Time) ([]Event, error)
}

type DateRepository interface {
	FindByDay(day string) (Date, error)
	Save(date Date) (uint32, error)
}

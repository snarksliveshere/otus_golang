package entity

type RecordRepository interface {
	FindById(id uint64) (Record, error)
	GetEventsByDay(dayFk uint32) ([]Record, error)
	Save(record Record) (uint64, error)
	Delete(record Record) error
	Edit(record Record) error
	GetEventsByDateInterval(from, till string) ([]Record, error)
	Show() []Record
}

type DateRepository interface {
	FindByDay(day string) (Date, error)
	Save(date Date) (uint32, error)
}

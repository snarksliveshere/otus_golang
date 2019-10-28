package entity

type RecordRepository interface {
	FindById(id uint64) (Record, error)
	Delete(record Record) error
	Edit(record Record) error
	Show() []Record
	Save(record Record) error
}

type DayRepository interface {
	AddRecordToDay(record Record, day Day) error
	ShowDayRecords(day Day) ([]Record, error)
	FindById(id uint64) (Day, error)
	Save(record Day) error
}

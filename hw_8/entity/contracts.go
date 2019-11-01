package entity

type RecordRepository interface {
	FindById(id uint64) (Record, error)
	Delete(record Record) error
	Edit(record Record) error
	Show() []Record
	Save(record Record) error
}

type DateRepository interface {
	AddRecordToDate(record Record, day Date) error
	ShowDayRecords(day Date) ([]Record, error)
	FindById(id uint64) (Date, error)
}

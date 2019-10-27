package entity

type Record struct {
	Id          uint64
	Title       string
	Description string
}

type RecordRepository interface {
	FindById(id uint64) (Record, error)
	Delete(record Record) error
	Edit(record Record) error
	Show() []Record
	Store(record Record) (Record, error)
}

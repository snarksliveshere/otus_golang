package entity

type Record struct {
	Id          uint64
	Title       string
	Description string
}

type RecordRepository interface {
	Add(record Record)
	FindById(id uint64) Record
	Delete(id uint64) bool
	Edit(id uint64, title, description string) bool
	Show() []Record
}

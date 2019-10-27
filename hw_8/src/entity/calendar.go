package entity

type Calendar struct {
	Year int16
	Month
	Week uint8
	Day
}

type Month struct {
	Name      string
	Num       uint8
	NumOfDays uint8
}

type Day struct {
	Id      uint64
	Records []Record
	Num     uint16
}

type DayRepository interface {
	AddRecordToDay(record Record, day Day) error
	ShowDayRecords(day Day) ([]Record, error)
	FindById(id uint64) (Day, error)
	Store(record Record) error
}

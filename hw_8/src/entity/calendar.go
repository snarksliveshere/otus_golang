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
	Records []Record
	Num     uint16
}

type DayRepository interface {
	ShowDayRecords(day Day) []Record
	AddRecordToDay(record Record) bool
}

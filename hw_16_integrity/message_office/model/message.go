package model

type Message struct {
	TableName struct{} `sql:"calendar.message"`
	Id        uint64
	Status    string `sql:"status, notnull"`
	Msg       string `sql:"msg"`
}

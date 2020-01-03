package pg_models

type Calendar struct {
	TableName     struct{} `sql:"calendar.calendar"`
	Id            uint32
	Date          string `sql:"date,notnull,unique"`
	Description   string `sql:"description"`
	IsCelebration bool   `sql:"is_celebration,notnull"`
}

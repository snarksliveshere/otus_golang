package mem_repository

import (
	"github.com/snarskliveshere/otus_golang/hw_8/src/entity"
)

func (d *DayRepo) AddRecordToDay(record entity.Record, day entity.Day) error {
	d.handler.Execute("add record to day")
	return nil
}

func (d *DayRepo) ShowDayRecords(day entity.Day) ([]entity.Record, error) {
	d.handler.Execute("show day record")
	return []entity.Record{}, nil
}

func (d *DayRepo) FindById(id uint64) (entity.Day, error) {
	// pg-go
	//model := model.Day{}
	//_, err := app.Db().Model(&model).
	//	Where("id = ?", id).
	//	Select()
	day := entity.Day{
		Id:      id,
		Records: []entity.Record{},
		Num:     2,
	}
	d.handler.Execute("find by id")
	return day, nil
}

func (d *DayRepo) Save(record entity.Day) error {
	d.handler.Execute("save")
	return nil
}

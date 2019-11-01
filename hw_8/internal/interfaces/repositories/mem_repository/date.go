package mem_repository

import (
	"github.com/snarskliveshere/otus_golang/hw_8/entity"
)

func (d *DateRepo) AddRecordToDate(record entity.Record, day entity.Date) error {
	d.handler.Execute("add record to day")
	return nil
}

func (d *DateRepo) ShowDayRecords(day entity.Date) ([]entity.Record, error) {
	d.handler.Execute("show day record")
	return []entity.Record{}, nil
}

func (d *DateRepo) FindById(id uint64) (entity.Date, error) {
	day := entity.Date{
		Id:      id,
		Records: []entity.Record{},
	}
	d.handler.Execute("find by id")
	return day, nil
}

func (d *DateRepo) Save(record entity.Date) error {
	d.handler.Execute("save")
	return nil
}

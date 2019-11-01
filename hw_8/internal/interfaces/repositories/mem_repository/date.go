package mem_repository

import (
	"github.com/snarskliveshere/otus_golang/hw_8/entity"
	"time"
)

func (d *DateRepo) AddRecordToDate(record entity.Record, day entity.Date) error {
	d.handler.Execute("add record to day")
	return nil
}

func (d *DateRepo) ShowDayRecords(day entity.Date) ([]entity.Record, error) {
	d.handler.Execute("show day record")
	return []entity.Record{}, nil
}

func (d *DateRepo) GetDateFromString(date string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, date)
	if err != nil {
		d.logger.Log("Wrong incoming day pattern")
	}
	return t, nil
}

func (d *DateRepo) FindByDay(date time.Time) (entity.Date, error) {
	day := entity.Date{
		Day:     date,
		Records: []entity.Record{},
	}

	d.handler.Execute("find by day")
	return day, nil
}

func (d *DateRepo) Save(record entity.Date) error {

	d.handler.Execute("save")
	return nil
}

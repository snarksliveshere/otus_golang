package mem_repository

import (
	"github.com/snarskliveshere/otus_golang/hw_8/entity"
	"time"
)

func (d *DateRepo) AddEventToDate(event entity.Event, day entity.Date) error {
	d.handler.Execute("add event to day")
	return nil
}

func (d *DateRepo) ShowDayEvents(day entity.Date) ([]entity.Event, error) {
	d.handler.Execute("show day event")
	return nil, nil
}

func (d *DateRepo) GetDateFromString(date string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, date)
	if err != nil {
		d.logger.Log().Info("Wrong incoming day pattern")
		return t, err
	}
	return t, nil
}

func (d *DateRepo) FindByDay(date time.Time) (entity.Date, error) {
	day := entity.Date{
		Day:    date,
		Events: []entity.Event{},
	}

	d.handler.Execute("find by day")
	return day, nil
}

func (d *DateRepo) Save(event entity.Date) error {

	d.handler.Execute("save")
	return nil
}

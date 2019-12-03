package mem_repository

import (
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/config"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/entity"
	"sync"
	"time"
)

var (
	calendar     *entity.Calendar
	calendarOnce sync.Once
)

func createCalendar() *entity.Calendar {
	calendarOnce.Do(func() {
		calendar = &entity.Calendar{}
	})
	return calendar
}

func (d *DateRepo) GetCalendar() *entity.Calendar {
	return createCalendar()
}

func (d *DateRepo) AddDateToCalendar(day entity.Date) error {
	calendar := d.GetCalendar()
	calendar.Dates = append(calendar.Dates, &day)
	d.handler.Execute("add date to calendar")
	return nil
}

func (d *DateRepo) AddEventToDate(event entity.Event, day *entity.Date) error {
	day.Events = append(day.Events, event)
	d.handler.Execute("add event to day")
	return nil
}

func (d *DateRepo) ShowDayEvents(day *entity.Date) ([]entity.Event, error) {
	d.handler.Execute("show day event")
	return nil, nil
}

func (d *DateRepo) GetDateFromString(date string) (time.Time, error) {
	t, err := time.Parse(config.TimeLayout, date)
	if err != nil {
		d.logger.Info("Wrong incoming day pattern")
		return t, err
	}
	return t, nil
}

func (d *DateRepo) FindByDay(date time.Time, calendar *entity.Calendar) (*entity.Date, error) {
	day := entity.Date{
		Day:    date,
		Events: []entity.Event{},
	}
	if calendar.Dates == nil {
		calendar.Dates = []*entity.Date{}
		calendar.Dates = append(calendar.Dates, &day)
		return &day, nil
	}

	var isDateInCalendarIndex int
	var isDateInCalendarBool bool
	for i, z := range calendar.Dates {
		fmt.Println(z.Day.Format(config.TimeLayout), date.Format(config.TimeLayout), z.Day.Format(config.TimeLayout) == date.Format(config.TimeLayout))
		if z.Day.Format(config.TimeLayout) == date.Format(config.TimeLayout) {
			isDateInCalendarIndex = i
			isDateInCalendarBool = true
		}
	}
	if isDateInCalendarBool {
		d.handler.Execute("day exist")
		return calendar.Dates[isDateInCalendarIndex], nil
	}
	calendar.Dates = append(calendar.Dates, &day)
	d.handler.Execute("add new day")
	return &day, nil
}

func (d *DateRepo) Save(event entity.Date) error {

	d.handler.Execute("save")
	return nil
}

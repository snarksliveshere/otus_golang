package mem_repository

import (
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/config"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/entity"
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

func (d *DateRepo) AddRecordToDate(record entity.Record, day *entity.Date) error {
	d.handler.Execute("add record to day")
	// TODO: EMPTY RECORDS ON DATE!!!
	fmt.Printf("\nBEFORE date!! %#v \n records !! %#v \n", day, record)
	day.Records = append(day.Records, record)
	fmt.Printf("\nAFTER date!! %#v \n records !! %#v \n", day, record)
	return nil
}

func (d *DateRepo) ShowDayRecords(day entity.Date) ([]entity.Record, error) {
	d.handler.Execute("show day record")
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

//func (d *DateRepo) FindByDay(date time.Time) (entity.Date, error) {
//	calendar := d.GetCalendar()
//	var isDateInCalendarIndex int
//	for i, z := range calendar.Dates {
//		if z.Day.Format(config.TimeLayout) == date.Format(config.TimeLayout) {
//			isDateInCalendarIndex = i
//		}
//	}
//	if isDateInCalendarIndex != 0 {
//		return calendar.Dates[isDateInCalendarIndex], nil
//	}
//
//	day := entity.Date{
//		Day:     date,
//		Records: []entity.Record{},
//	}
//	calendar.Dates = append(calendar.Dates, day)
//
//	d.handler.Execute("find by day")
//	return day, nil
//}

func (d *DateRepo) FindByDay(date time.Time, calendar *entity.Calendar) (entity.Date, error) {
	if calendar.Dates == nil {
		day := entity.Date{
			Day:     date,
			Records: []entity.Record{},
		}
		calendar.Dates = []*entity.Date{}
		calendar.Dates = append(calendar.Dates, &day)
		return day, nil
	}

	var isDateInCalendarIndex int
	for i, z := range calendar.Dates {
		if z.Day.Format(config.TimeLayout) == date.Format(config.TimeLayout) {
			isDateInCalendarIndex = i
		}
	}
	var day entity.Date
	if isDateInCalendarIndex != 0 {
		ds := calendar.Dates
		d.handler.Execute("day exist")
		return *ds[isDateInCalendarIndex], nil
	}
	calendar.Dates = append(calendar.Dates, &day)
	d.handler.Execute("add new day")
	return day, nil
}

func (d *DateRepo) Save(record entity.Date) error {

	d.handler.Execute("save")
	return nil
}

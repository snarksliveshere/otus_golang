package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/entity"
	"time"
)

func (act *Actions) AddRecordToDate(rec entity.Record, date time.Time) (entity.Date, error) {
	day, err := act.getDate(date)
	if err != nil {
		act.Logger.Info(err.Error())
		return entity.Date{}, err
	}
	err = act.DateRepository.AddRecordToDate(rec, &day)
	if err != nil {
		act.Logger.Info("An error occurred while add day records")
		return entity.Date{}, err
	}

	return day, nil
}

func (act *Actions) AddDateToCalendar(date entity.Date) error {
	err := act.DateRepository.AddDateToCalendar(date)
	if err != nil {
		act.Logger.Info("An error occurred while add day records")
		return err
	}
	return nil
}

func (act *Actions) ShowDayRecords(dateStr string) ([]entity.Record, error) {
	day, err := act.DateRepository.GetDateFromString(dateStr)
	if err != nil {
		act.Logger.Info("An error occurred while get day in time.Time")
		return []entity.Record{}, err
	}
	date, err := act.DateRepository.FindByDay(day)
	records, err := act.DateRepository.ShowDayRecords(date)
	if err != nil {
		act.Logger.Info("An error occurred while show day records")
		return []entity.Record{}, err
	}
	return records, nil
}

func (act *Actions) getDate(date time.Time) (entity.Date, error) {
	entityDate, err := act.DateRepository.FindByDay(date)
	if err != nil {
		act.Logger.Info("An error occurred while get day")
		return entity.Date{}, err
	}
	return entityDate, nil
}

func (act *Actions) returnDate(day time.Time) entity.Date {
	date := entity.Date{
		Day:     day,
		Records: nil,
	}
	return date
}

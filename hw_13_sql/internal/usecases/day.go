package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_13_sql/config"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
	"time"
)

func (act *Actions) FindByDay(day string) (entity.Date, error) {
	date, err := act.DateRepository.FindByDay(day)
	if err != nil {
		return entity.Date{}, err
	}
	return date, nil
}

func (act *Actions) AddRecordToDate(rec entity.Record, date *entity.Date) error {
	err := act.DateRepository.AddRecordToDate(rec, date)
	if err != nil {
		act.Logger.Info("An error occurred while add day records")
		return err
	}

	return nil
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
	_, err := act.DateRepository.GetDateFromString(dateStr)
	if err != nil {
		act.Logger.Info("An error occurred while get day in time.Time")
		return []entity.Record{}, err
	}
	date, err := act.DateRepository.FindByDay(config.TimeLayout)
	records, err := act.DateRepository.ShowDayRecords(&date)
	if err != nil {
		act.Logger.Info("An error occurred while show day records")
		return []entity.Record{}, err
	}
	return records, nil
}

func (act *Actions) getDate(date time.Time) (entity.Date, error) {
	entityDate, err := act.DateRepository.FindByDay(config.TimeLayout)
	if err != nil {
		act.Logger.Info("An error occurred while get day")
		return entity.Date{}, err
	}
	return entityDate, nil
}

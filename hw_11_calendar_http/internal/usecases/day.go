package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/entity"
	"time"
)

func (act *Actions) AddRecordToDate(recId int, dateStr string) error {
	date, err := act.getDate(dateStr)
	if err != nil {
		act.Logger.Log(err.Error())
		return err
	}
	rec, err := act.getRecordById(recId)
	if err != nil {
		act.Logger.Log(err.Error())
		return err
	}
	err = act.DateRepository.AddRecordToDate(rec, date)
	if err != nil {
		act.Logger.Log("An error occurred while add day records")
		return err
	}

	return nil
}

func (act *Actions) ShowDayRecords(dateStr string) ([]entity.Record, error) {
	day, err := act.DateRepository.GetDateFromString(dateStr)
	if err != nil {
		act.Logger.Log("An error occurred while get day in time.Time")
		return []entity.Record{}, err
	}
	date, err := act.DateRepository.FindByDay(day)
	records, err := act.DateRepository.ShowDayRecords(date)
	if err != nil {
		act.Logger.Log("An error occurred while show day records")
		return []entity.Record{}, err
	}
	return records, nil
}

func (act *Actions) getDate(dateStr string) (entity.Date, error) {
	day, err := act.DateRepository.GetDateFromString(dateStr)
	if err != nil {
		act.Logger.Log("An error occurred while get day in time.Time")
		return entity.Date{}, err
	}
	date, err := act.DateRepository.FindByDay(day)
	if err != nil {
		act.Logger.Log("An error occurred while get day")
		return entity.Date{}, err
	}
	return date, nil
}

func (act *Actions) returnDate(day time.Time) entity.Date {
	date := entity.Date{
		Day:     day,
		Records: nil,
	}
	return date
}

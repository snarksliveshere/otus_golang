package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_8/entity"
	"time"
)

func (act *Actions) AddEventToDate(recId int, dateStr string) error {
	date, err := act.getDate(dateStr)
	if err != nil {
		act.Logger.Log(err.Error())
		return err
	}
	rec, err := act.getEventById(recId)
	if err != nil {
		act.Logger.Log(err.Error())
		return err
	}
	err = act.DateRepository.AddEventToDate(rec, date)
	if err != nil {
		act.Logger.Log("An error occurred while add day events")
		return err
	}

	return nil
}

func (act *Actions) ShowDayEvents(dateStr string) ([]entity.Event, error) {
	day, err := act.DateRepository.GetDateFromString(dateStr)
	if err != nil {
		act.Logger.Log("An error occurred while get day in time.Time")
		return []entity.Event{}, err
	}
	date, err := act.DateRepository.FindByDay(day)
	events, err := act.DateRepository.ShowDayEvents(date)
	if err != nil {
		act.Logger.Log("An error occurred while show day events")
		return []entity.Event{}, err
	}
	return events, nil
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
		Day:    day,
		Events: nil,
	}
	return date
}

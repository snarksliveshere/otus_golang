package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/entity"
	"time"
)

func (act *Actions) AddEventToDate(rec entity.Event, date *entity.Date) error {
	err := act.DateRepository.AddEventToDate(rec, date)
	if err != nil {
		act.Logger.Info("An error occurred while add day events")
		return err
	}

	return nil
}

func (act *Actions) AddDateToCalendar(date entity.Date) error {
	err := act.DateRepository.AddDateToCalendar(date)
	if err != nil {
		act.Logger.Info("An error occurred while add day events")
		return err
	}
	return nil
}

func (act *Actions) ShowDayEvents(dateStr string) ([]entity.Event, error) {
	day, err := act.DateRepository.GetDateFromString(dateStr)
	if err != nil {
		act.Logger.Info("An error occurred while get day in time.Time")
		return []entity.Event{}, err
	}
	date, err := act.DateRepository.FindByDay(day, &entity.Calendar{})
	events, err := act.DateRepository.ShowDayEvents(date)
	if err != nil {
		act.Logger.Info("An error occurred while show day events")
		return []entity.Event{}, err
	}
	return events, nil
}

//func (act *Actions) getDate(date time.Time) (entity.Date, error) {
//	entityDate, err := act.DateRepository.FindByDay(date)
//	if err != nil {
//		act.Logger.Info("An error occurred while get day")
//		return entity.Date{}, err
//	}
//	return entityDate, nil
//}

func (act *Actions) returnDate(day time.Time) entity.Date {
	date := entity.Date{
		Day:    day,
		Events: nil,
	}
	return date
}

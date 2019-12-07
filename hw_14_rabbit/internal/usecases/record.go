package usecases

import (
	"errors"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/config"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/entity"
	"strings"
	"time"
)

func (act *Actions) CreateEvent(title, description string, t time.Time) (uint64, error) {
	date := t.Format(config.TimeLayout)
	day, err := act.DateRepository.FindByDay(date)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return 0, err
	}
	var dateId uint32
	if day.Id == 0 {
		newDateId, err := act.CreateEmptyDate(date)
		if err != nil {
			return 0, err
		}
		dateId = newDateId
	} else {
		dateId = day.Id
	}
	recId, err := act.AddEvent(title, description, dateId, t)

	return recId, nil
}

func (act *Actions) DeleteEventById(id uint64) error {
	rec, err := act.EventRepository.FindById(id)
	if err != nil {
		return err
	}
	err = act.EventRepository.Delete(rec)
	if err != nil {
		return err
	}
	return nil
}

func (act *Actions) UpdateEventById(id uint64, title, description string) error {
	rec, err := act.EventRepository.FindById(id)
	if err != nil {
		return err
	}
	rec.Title = title
	rec.Description = description

	err = act.EventRepository.Edit(rec)
	if err != nil {
		return err
	}
	return nil
}

func (act *Actions) AddEvent(title, description string, dateFk uint32, t time.Time) (id uint64, err error) {
	rec := entity.Event{
		Title:       title,
		Description: description,
		Time:        t,
		DateFk:      dateFk,
	}

	recId, err := act.EventRepository.Save(rec)
	if err != nil {
		act.Logger.Info("An error occurred while event added")
		return 0, err
	}
	act.Logger.Info("Event added successfully")

	return recId, nil
}

func (act *Actions) GetEventsByDay(date string) ([]entity.Event, error) {
	day, err := act.DateRepository.FindByDay(date)
	if err != nil {
		act.Logger.Info("An error occurred while event added")
		return []entity.Event{}, err
	}
	if day.Id == 0 {
		act.Logger.Info("An empty ady")
		return []entity.Event{}, errors.New("i cant find this day")
	}
	events, err := act.EventRepository.GetEventsByDay(day.Id)
	if err != nil {
		act.Logger.Info("An error occurred while event added")
		return []entity.Event{}, err
	}

	return events, nil
}

func (act *Actions) EditEvent(id uint64) error {
	rec, err := act.getEventById(id)
	if err != nil {
		act.Logger.Info(err.Error())
		return err
	}
	err = act.EventRepository.Edit(rec)
	if err != nil {
		act.Logger.Info("An error occurred while event updating")
		return err
	}
	return nil
}

func (act *Actions) DeleteEvent(id uint64) error {
	rec, err := act.getEventById(id)
	if err != nil {
		act.Logger.Info(err.Error())
		return err
	}
	err = act.EventRepository.Delete(rec)
	if err != nil {
		act.Logger.Info("An error occurred while event deleting")
		return err
	}
	return nil
}

func (act *Actions) getEventById(id uint64) (entity.Event, error) {
	event, err := act.EventRepository.FindById(id)
	if err != nil {
		act.Logger.Info("An error occurred while get event")
		return entity.Event{}, err
	}
	return event, nil
}

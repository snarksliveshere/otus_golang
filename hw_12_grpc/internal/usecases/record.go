package usecases

import (
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/entity"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/internal/helpers"
)

func (act *Actions) AddEvent(title, description string) (rec entity.Event, err error) {
	id := helpers.MakeTimestampId()
	rec = entity.Event{
		Id:          id,
		Title:       title,
		Description: description,
	}

	err = act.EventRepository.Save(rec)
	if err != nil {
		act.Logger.Info("An error occurred while event added")
		return rec, err
	}
	act.Logger.Info("Event added successfully")

	return rec, nil
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

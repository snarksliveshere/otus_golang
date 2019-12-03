package usecases

import "github.com/snarskliveshere/otus_golang/hw_8/entity"

func (act *Actions) AddEvent(title, description string) error {
	rec := entity.Event{
		Id:          1,
		Title:       title,
		Description: description,
	}

	err := act.EventRepository.Save(rec)
	if err != nil {
		act.Logger.Log("An error occurred while event added")
		return err
	}
	act.Logger.Log("Event added successfully")

	return nil
}

func (act *Actions) EditEvent(id int) error {
	rec, err := act.getEventById(id)
	if err != nil {
		act.Logger.Log(err.Error())
		return err
	}
	err = act.EventRepository.Edit(rec)
	if err != nil {
		act.Logger.Log("An error occurred while event updating")
		return err
	}
	return nil
}

func (act *Actions) DeleteEvent(id int) error {
	rec, err := act.getEventById(id)
	if err != nil {
		act.Logger.Log(err.Error())
		return err
	}
	err = act.EventRepository.Delete(rec)
	if err != nil {
		act.Logger.Log("An error occurred while event deleting")
		return err
	}
	return nil
}

func (act *Actions) getEventById(id int) (entity.Event, error) {
	event, err := act.EventRepository.FindById(uint64(id))
	if err != nil {
		act.Logger.Log("An error occurred while get event")
		return entity.Event{}, err
	}
	return event, nil
}

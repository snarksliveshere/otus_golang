package inmem

import (
	"errors"
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/config"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/entity"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/internal/interfaces/repositories/mem_repository"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/internal/usecases"
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/pkg"
	"time"
)

type Storage struct {
	actions  *usecases.Actions
	calendar *entity.Calendar
}

//func InMemFunc() {
//	handler := pkg.NewStorageHandler()
//	actions := new(usecases.Actions)
//	actions.Logger = new(pkg.Logger)
//	actions.DateRepository = mem_repository.GetDateRepo(handler)
//	actions.EventRepository = mem_repository.GetEventRepo(handler)
//}

func CreateStorageInstance(logger usecases.Logger) *Storage {
	handler := pkg.NewStorageHandler()
	actions := new(usecases.Actions)
	actions.Logger = logger
	actions.DateRepository = mem_repository.GetDateRepo(handler)
	actions.EventRepository = mem_repository.GetEventRepo(handler)
	c := new(entity.Calendar)
	return &Storage{actions: actions, calendar: c}
}

func (s *Storage) AddEvent(title, desc string, date time.Time) (entity.Event, *entity.Date, *entity.Calendar, error) {
	rec, err := s.actions.AddEvent(title, desc)
	if err != nil {
		return entity.Event{}, &entity.Date{}, s.calendar, err
	}
	day, err := s.actions.DateRepository.FindByDay(date, s.calendar)
	if err != nil {
		return rec, &entity.Date{}, s.calendar, err
	}
	err = s.actions.DateRepository.AddEventToDate(rec, day)
	if err != nil {
		return rec, &entity.Date{}, s.calendar, err
	}

	return rec, day, s.calendar, nil
}

func (s *Storage) FindEventById(id uint64) string {
	event, _ := s.actions.EventRepository.FindById(id)
	return fmt.Sprintf("resccc %#v", event)
}

func (s *Storage) DeleteEventById(id uint64) error {
	if s.calendar.Dates == nil {
		err := errors.New("there are no events in calendar yet")
		return err
	}
	var res bool
	for _, z := range s.calendar.Dates {
		for i, r := range z.Events {
			if r.Id == id {
				newEvents := removeEventFromSlice(z.Events, i)
				z.Events = append([]entity.Event(nil), newEvents...)
				res = true
			}
		}
	}
	if res {
		return nil
	} else {
		err := errors.New("i cant find event with this id to delete")
		return err
	}
}

func (s *Storage) UpdateEventById(recId uint64, date time.Time, title, description string) error {
	if s.calendar.Dates == nil {
		err := errors.New("there are no events in calendar yet")
		return err
	}
	var res bool
	for i, z := range s.calendar.Dates {
		if z.Day.Format(config.TimeLayout) == date.Format(config.TimeLayout) {
			for k, r := range z.Events {
				if r.Id == recId {
					updEvent(&s.calendar.Dates[i].Events[k], title, description)
					res = true
				}
			}
		}
	}

	if res {
		return nil
	} else {
		err := errors.New("i cant find event with this id to update")
		return err
	}
}

func updEvent(rec *entity.Event, title, desc string) {
	rec.Title = title
	rec.Description = desc
}

func removeEventFromSlice(events []entity.Event, i int) []entity.Event {
	events[i] = events[len(events)-1]
	return events[:len(events)-1]
}

func (s *Storage) GetEventsForDay(date time.Time) (*entity.Date, error) {
	day, err := s.actions.DateRepository.FindByDay(date, s.calendar)
	if err != nil {
		return &entity.Date{}, err
	}
	return day, nil
}

func (s *Storage) GetEventsForInterval(from, till time.Time) ([]entity.Event, error) {
	if s.calendar.Dates == nil {
		err := errors.New("there are no events in calendar yet")
		return nil, err
	}
	var res bool
	var events []entity.Event
	for _, z := range s.calendar.Dates {
		if z.Day.Format(config.TimeLayout) >= from.Format(config.TimeLayout) &&
			z.Day.Format(config.TimeLayout) <= till.Format(config.TimeLayout) {
			events = append(events, z.Events...)
			res = true
		}
	}

	if res {
		return events, nil
	} else {
		err := errors.New("i cant find events for this interval")
		return nil, err
	}
}

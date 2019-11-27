package pg_repository

import (
	"errors"
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/config"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/usecases"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/pkg"
	"time"
)

type Storage struct {
	actions  *usecases.Actions
	calendar *entity.Calendar
}

func CreateStorageInstance(logger usecases.Logger, conf *config.Config) *Storage {
	handler := pkg.NewStorageHandler()
	actions := new(usecases.Actions)
	actions.Logger = logger
	actions.DateRepository = GetDateRepo(handler)
	actions.RecordRepository = GetRecordRepo(handler)
	c := new(entity.Calendar)
	return &Storage{actions: actions, calendar: c}
}

func (s *Storage) AddRecord(title, desc string, date time.Time) (entity.Record, *entity.Date, *entity.Calendar, error) {
	rec, err := s.actions.AddRecord(title, desc)
	if err != nil {
		return entity.Record{}, &entity.Date{}, s.calendar, err
	}
	day, err := s.actions.DateRepository.FindByDay(date, s.calendar)
	if err != nil {
		return rec, &entity.Date{}, s.calendar, err
	}
	err = s.actions.DateRepository.AddRecordToDate(rec, day)
	if err != nil {
		return rec, &entity.Date{}, s.calendar, err
	}

	return rec, day, s.calendar, nil
}

func (s *Storage) FindRecordById(id uint64) (string, error) {
	record, err := s.actions.RecordRepository.FindById(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("resccc %#v", record), nil
}

func (s *Storage) DeleteRecordById(id uint64) error {
	if s.calendar.Dates == nil {
		return errors.New("there are no records in calendar yet")
	}
	var res bool
	for _, z := range s.calendar.Dates {
		for i, r := range z.Records {
			if r.Id == id {
				newRecords := removeRecordFromSlice(z.Records, i)
				z.Records = append(z.Records, newRecords...)
				res = true
			}
		}
	}
	if res {
		return nil
	} else {
		return errors.New("i cant find record with this id to delete")
	}
}

func (s *Storage) UpdateRecordById(recId uint64, date time.Time, title, description string) error {
	if s.calendar.Dates == nil {
		return errors.New("there are no records in calendar yet")
	}
	var res bool
	for i, z := range s.calendar.Dates {
		if z.Day.Format(config.TimeLayout) == date.Format(config.TimeLayout) {
			for k, r := range z.Records {
				if r.Id == recId {
					updRecord(&s.calendar.Dates[i].Records[k], title, description)
					res = true
				}
			}
		}
	}

	if res {
		return nil
	} else {
		err := errors.New("i cant find record with this id to update")
		return err
	}
}

func updRecord(rec *entity.Record, title, desc string) {
	rec.Title = title
	rec.Description = desc
}

func removeRecordFromSlice(records []entity.Record, i int) []entity.Record {
	records[i] = records[len(records)-1]
	return records[:len(records)-1]
}

func (s *Storage) GetEventsForDay(date time.Time) (*entity.Date, error) {
	day, err := s.actions.DateRepository.FindByDay(date, s.calendar)
	if err != nil {
		return nil, err
	}
	return day, nil
}

func (s *Storage) GetEventsForInterval(from, till time.Time) ([]entity.Record, error) {
	if s.calendar.Dates == nil {
		err := errors.New("there are no records in calendar yet")
		return nil, err
	}
	var res bool
	var records []entity.Record
	for _, z := range s.calendar.Dates {
		if z.Day.Format(config.TimeLayout) >= from.Format(config.TimeLayout) &&
			z.Day.Format(config.TimeLayout) <= till.Format(config.TimeLayout) {
			records = append(records, z.Records...)
			res = true
		}
	}

	if res {
		return records, nil
	} else {
		err := errors.New("i cant find records for this interval")
		return nil, err
	}
}

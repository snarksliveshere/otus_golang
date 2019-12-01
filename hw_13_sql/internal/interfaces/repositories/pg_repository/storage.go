package pg_repository

import (
	"errors"
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/config"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/entity"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/usecases"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/tools/db"
	"time"
)

type Storage struct {
	Actions  *usecases.Actions
	calendar *entity.Calendar
}

func CreateStorageInstance(logger usecases.Logger, conf *config.Config) *Storage {
	dbHandler := db.CreatePgConn(conf, logger)
	actions := new(usecases.Actions)
	actions.Logger = logger
	actions.DateRepository = GetDateRepo(dbHandler)
	actions.RecordRepository = GetRecordRepo(dbHandler)
	c := new(entity.Calendar)
	return &Storage{Actions: actions, calendar: c}
}

func (s *Storage) FindByDay(date string) (entity.Date, error) {
	return s.Actions.DateRepository.FindByDay(date)
}

func (s *Storage) GetEventsForDay(date string) ([]entity.Record, error) {
	records, err := s.Actions.GetEventsByDay(date)

	if err != nil {
		return []entity.Record{}, err
	}

	return records, nil
}

func (s *Storage) AddRecord(title, desc string, date time.Time) (entity.Record, *entity.Date, *entity.Calendar, error) {
	return entity.Record{}, &entity.Date{}, s.calendar, nil
}

func (s *Storage) FindRecordById(id uint64) string {
	record, _ := s.Actions.RecordRepository.FindById(id)
	return fmt.Sprintf("resccc %#v", record)
}

func (s *Storage) DeleteRecordById(id uint64) error {
	if s.calendar.Dates == nil {
		err := errors.New("there are no records in calendar yet")
		return err
	}
	var res bool
	for _, z := range s.calendar.Dates {
		for i, r := range z.Records {
			if r.Id == id {
				newRecords := removeRecordFromSlice(z.Records, i)
				z.Records = append([]entity.Record(nil), newRecords...)
				res = true
			}
		}
	}
	if res {
		return nil
	} else {
		err := errors.New("i cant find record with this id to delete")
		return err
	}
}

func (s *Storage) UpdateRecordById(recId uint64, date time.Time, title, description string) error {
	if s.calendar.Dates == nil {
		err := errors.New("there are no records in calendar yet")
		return err
	}
	var res bool
	for i, z := range s.calendar.Dates {
		if z.Day == date.Format(config.TimeLayout) {
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

func (s *Storage) GetEventsForInterval(from, till time.Time) ([]entity.Record, error) {
	if s.calendar.Dates == nil {
		err := errors.New("there are no records in calendar yet")
		return nil, err
	}
	var res bool
	var records []entity.Record
	for _, z := range s.calendar.Dates {
		if z.Day >= from.Format(config.TimeLayout) &&
			z.Day <= till.Format(config.TimeLayout) {
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

package inmem

import (
	"errors"
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/entity"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/interfaces/repositories/mem_repository"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/usecases"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/pkg"
	"time"
)

type Storage struct {
	actions *usecases.Actions
}

//func InMemFunc() {
//	handler := pkg.NewStorageHandler()
//	actions := new(usecases.Actions)
//	actions.Logger = new(pkg.Logger)
//	actions.DateRepository = mem_repository.GetDateRepo(handler)
//	actions.RecordRepository = mem_repository.GetRecordRepo(handler)
//}

func CreateStorageInstance(logger usecases.Logger) *Storage {
	handler := pkg.NewStorageHandler()
	actions := new(usecases.Actions)
	actions.Logger = logger
	actions.DateRepository = mem_repository.GetDateRepo(handler)
	actions.RecordRepository = mem_repository.GetRecordRepo(handler)

	return &Storage{actions: actions}
}

func (s *Storage) AddRecord(title, desc string, date time.Time) (entity.Record, entity.Date, error) {
	rec, err := s.actions.AddRecord(title, desc)
	if err != nil {
		return entity.Record{}, entity.Date{}, err
	}
	day, err := s.actions.AddRecordToDate(rec, date)
	if err != nil {
		return rec, entity.Date{}, err
	}

	return rec, day, nil
}

func (s *Storage) FindRecordById(id uint64) string {
	record, _ := s.actions.RecordRepository.FindById(id)
	return fmt.Sprintf("resccc %#v", record)
}

func (s *Storage) DeleteRecordById(id uint64) error {
	c := s.actions.DateRepository.GetCalendar()
	var res bool
	for _, z := range c.Dates {
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
		err := errors.New("i cant find record with this id")
		return err
	}
}

func removeRecordFromSlice(records []entity.Record, i int) []entity.Record {
	records[i] = records[len(records)-1]
	return records[:len(records)-1]
}

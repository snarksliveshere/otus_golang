package inmem

import (
	"errors"
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/config"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/entity"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/interfaces/repositories/mem_repository"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/usecases"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/pkg"
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
//	actions.RecordRepository = mem_repository.GetRecordRepo(handler)
//}

func CreateStorageInstance(logger usecases.Logger) *Storage {
	handler := pkg.NewStorageHandler()
	actions := new(usecases.Actions)
	actions.Logger = logger
	actions.DateRepository = mem_repository.GetDateRepo(handler)
	actions.RecordRepository = mem_repository.GetRecordRepo(handler)
	//c := actions.DateRepository.GetCalendar()
	c := new(entity.Calendar)
	return &Storage{actions: actions, calendar: c}
}

func (s *Storage) AddRecord(title, desc string, date time.Time) (entity.Record, entity.Date, error) {
	//c := s.calendar
	fmt.Printf("\nc before ol1 %#v\n", s.calendar)
	rec, err := s.actions.AddRecord(title, desc)
	if err != nil {
		return entity.Record{}, entity.Date{}, err
	}
	day, _ := s.actions.DateRepository.FindByDay(date, s.calendar)
	fmt.Printf("\nday with records: %#v\n", day)
	fmt.Printf("\ncalendar with day: %#v\n", s.calendar)
	err = s.actions.DateRepository.AddRecordToDate(rec, &day)
	// TODO: no date in calendar
	fmt.Printf("\ncalendar with day with records: %#v\n", s.calendar)
	fmt.Printf("\ndate::: %#v\n", day)
	if err != nil {
		return rec, entity.Date{}, err
	}
	fmt.Printf("\nc after ol2 %#v\n", s.calendar)

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
		err := errors.New("i cant find record with this id to delete")
		return err
	}
}

func (s *Storage) UpdateRecordById(recId uint64, date time.Time, title, description string) error {
	c := s.actions.DateRepository.GetCalendar()
	fmt.Printf("c before: %#v\n", c)
	var res bool
	for _, z := range c.Dates {
		if z.Day.Format(config.TimeLayout) == date.Format(config.TimeLayout) {
			for _, r := range z.Records {
				if r.Id == recId {
					r.Title = title
					r.Description = description
					res = true
				}
			}
		}
	}
	fmt.Printf("c after: %#v\n", c)
	if res {
		return nil
	} else {
		err := errors.New("i cant find record with this id to update")
		return err
	}
}

func removeRecordFromSlice(records []entity.Record, i int) []entity.Record {
	records[i] = records[len(records)-1]
	return records[:len(records)-1]
}

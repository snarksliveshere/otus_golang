package inmem

import (
	"fmt"
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

func (s *Storage) AddRecord(title, desc string, date time.Time) error {
	rec, err := s.actions.AddRecord(title, desc)
	if err != nil {
		return err
	}
	err = s.actions.AddRecordToDate(int(rec.Id), date)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) FindRecordById(id uint64) string {
	record, _ := s.actions.RecordRepository.FindById(id)
	return fmt.Sprintf("resccc %#v", record)
}

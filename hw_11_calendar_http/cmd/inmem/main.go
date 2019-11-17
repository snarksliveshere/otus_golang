package inmem

import (
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/interfaces/repositories/mem_repository"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/usecases"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/pkg"
)

type Storage struct {
	actions *usecases.Actions
}

func InMemFunc() {
	handler := pkg.NewStorageHandler()
	actions := new(usecases.Actions)
	actions.Logger = new(pkg.Logger)
	actions.DateRepository = mem_repository.GetDateRepo(handler)
	actions.RecordRepository = mem_repository.GetRecordRepo(handler)
	err := actions.AddRecord("title1", "descr1")
	if err != nil {
		fmt.Println(err.Error())
	}
	record, err := actions.RecordRepository.FindById(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(record)
}

func CreateStorageInstance(logger usecases.Logger) *Storage {
	handler := pkg.NewStorageHandler()
	actions := new(usecases.Actions)
	actions.Logger = logger
	actions.DateRepository = mem_repository.GetDateRepo(handler)
	actions.RecordRepository = mem_repository.GetRecordRepo(handler)

	return &Storage{actions: actions}
}

func (s *Storage) AddRecord(title, desc string) error {
	err := s.actions.AddRecord(title, desc)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) FindRecordById(id uint64) string {
	record, _ := s.actions.RecordRepository.FindById(id)
	return fmt.Sprintf("resccc %#v", record)
}

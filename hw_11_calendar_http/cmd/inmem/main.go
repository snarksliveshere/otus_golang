package inmem

import (
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/interfaces/repositories/mem_repository"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/internal/usecases"
	"github.com/snarskliveshere/otus_golang/hw_11_calendar_http/pkg"
)

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

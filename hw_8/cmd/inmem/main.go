package inmem

import (
	"fmt"
	"github.com/snarskliveshere/otus_golang/hw_8/internal/interfaces/repositories/mem_repository"
	"github.com/snarskliveshere/otus_golang/hw_8/internal/usecases"
	"github.com/snarskliveshere/otus_golang/hw_8/pkg"
)

func InMemFunc() {
	handler := pkg.NewStorageHandler()
	actions := new(usecases.Actions)
	actions.Logger = new(pkg.Logger)
	actions.DateRepository = mem_repository.GetDateRepo(handler)
	actions.EventRepository = mem_repository.GetEventRepo(handler)
	err := actions.AddEvent("title1", "descr1")
	if err != nil {
		fmt.Println(err.Error())
	}
	event, err := actions.EventRepository.FindById(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(event)
}

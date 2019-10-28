package main

import (
	"github.com/snarskliveshere/otus_golang/hw_8/src/drivers"
	"github.com/snarskliveshere/otus_golang/hw_8/src/interfaces/repositories/mem_repository"
	"github.com/snarskliveshere/otus_golang/hw_8/src/usecases"
)

func main() {
	handler := drivers.NewStorageHandler()
	//repo := mem_repository.CreateRepo(handler)
	//r := repo["dayRepo"]

	actions := new(usecases.Actions)
	actions.Logger = new(drivers.Logger)
	actions.DayRepository = mem_repository.GetDayRepo(handler)
	actions.RecordRepository = mem_repository.GetRecordRepo(handler)
}

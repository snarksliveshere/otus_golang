package main

import (
	"github.com/snarskliveshere/otus_golang/hw_8/internal/interfaces/repositories/mem_repository"
	"github.com/snarskliveshere/otus_golang/hw_8/internal/usecases"
	"github.com/snarskliveshere/otus_golang/hw_8/pkg"
)

func main() {
	handler := pkg.NewStorageHandler()
	actions := new(usecases.Actions)
	actions.Logger = new(pkg.Logger)
	actions.DateRepository = mem_repository.GetDateRepo(handler)
	actions.RecordRepository = mem_repository.GetRecordRepo(handler)
}

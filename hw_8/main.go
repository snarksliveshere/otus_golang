package main

import (
	drivers "github.com/snarskliveshere/otus_golang/hw_8/src/drivers"
	"github.com/snarskliveshere/otus_golang/hw_8/src/usecases"
)

func main() {
	dbHandler := drivers.NewStorageHandler()

	handlers := make(map[string]interfaces.DbHandler)
	handlers["DbUserRepo"] = dbHandler
	handlers["DbCustomerRepo"] = dbHandler
	handlers["DbItemRepo"] = dbHandler
	handlers["DbOrderRepo"] = dbHandler

	actions := new(usecases.Actions)
	actions.Logger = new(drivers.Logger)
	orderInteractor.UserRepository = interfaces.NewDbUserRepo(handlers)
	orderInteractor.ItemRepository = interfaces.NewDbItemRepo(handlers)
	orderInteractor.OrderRepository = interfaces.NewDbOrderRepo(handlers)

}

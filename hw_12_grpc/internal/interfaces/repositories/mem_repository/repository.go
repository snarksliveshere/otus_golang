package mem_repository

import (
	"github.com/snarskliveshere/otus_golang/hw_12_grpc/pkg"
)

type DateRepo struct {
	*Repo
}

type EventRepo struct {
	*Repo
}

type Repo struct {
	handler *pkg.InMemStorage
	logger  pkg.Logger
}

func (r *Repo) Execute(str string) {
	r.handler.Execute("execute")
}

//func CreateRepo(handler *drivers.InMemStorage) map[string]interface{} {
//	repo := new(Repo)
//	repo.handler = handler
//	m := make(map[string]interface{},2)
//
//	m["dayRepo"] = DayRepo{repo}
//	m["eventRepo"] = EventRepo{repo}
//
//	return m
//}

func GetDateRepo(handler *pkg.InMemStorage) *DateRepo {
	repo := new(Repo)
	repo.handler = handler
	return &DateRepo{repo}
}

func GetEventRepo(handler *pkg.InMemStorage) *EventRepo {
	repo := new(Repo)
	repo.handler = handler
	return &EventRepo{repo}
}

//func CreateRepo(handler *drivers.InMemStorage) *Repo {
//	return &Repo{handler: handler}
//}

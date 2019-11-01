package mem_repository

import (
	"github.com/snarskliveshere/otus_golang/hw_8/pkg"
)

type DateRepo struct {
	*Repo
}

type RecordRepo struct {
	*Repo
}

type Repo struct {
	handler *pkg.InMemStorage
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
//	m["recordRepo"] = RecordRepo{repo}
//
//	return m
//}

func GetDateRepo(handler *pkg.InMemStorage) *DateRepo {
	repo := new(Repo)
	repo.handler = handler
	return &DateRepo{repo}
}

func GetRecordRepo(handler *pkg.InMemStorage) *RecordRepo {
	repo := new(Repo)
	repo.handler = handler
	return &RecordRepo{repo}
}

//func CreateRepo(handler *drivers.InMemStorage) *Repo {
//	return &Repo{handler: handler}
//}

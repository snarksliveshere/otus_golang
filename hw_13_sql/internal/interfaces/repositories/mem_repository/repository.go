package mem_repository

import (
	"github.com/snarskliveshere/otus_golang/hw_13_sql/tools/db"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/tools/logger"
)

type DateRepo struct {
	*Repo
}

type RecordRepo struct {
	*Repo
}

type Repo struct {
	handler *db.InMemStorage
	logger  logger.Logger
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

func GetDateRepo(handler *db.InMemStorage) *DateRepo {
	repo := new(Repo)
	repo.handler = handler
	return &DateRepo{repo}
}

func GetRecordRepo(handler *db.InMemStorage) *RecordRepo {
	repo := new(Repo)
	repo.handler = handler
	return &RecordRepo{repo}
}

//func CreateRepo(handler *drivers.InMemStorage) *Repo {
//	return &Repo{handler: handler}
//}

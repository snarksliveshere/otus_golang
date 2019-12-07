package pg_repository

import (
	"github.com/go-pg/pg"
	pg_models2 "github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/internal/interfaces/repositories/pg_repository/pg_models"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/pkg/logger/logrus"
)

type DateRepo struct {
	*Repo
	row  *pg_models2.Calendar
	rows []*pg_models2.Calendar
}

type EventRepo struct {
	*Repo
	row  *pg_models2.Event
	rows []*pg_models2.Event
}

type Repo struct {
	db     *pg.DB
	logger logrus.Logger
}

func GetDateRepo(db *pg.DB) *DateRepo {
	repo := new(Repo)
	repo.db = db
	return &DateRepo{
		Repo: repo,
		row:  new(pg_models2.Calendar),
		rows: []*pg_models2.Calendar{},
	}
}

func GetEventRepo(db *pg.DB) *EventRepo {
	repo := new(Repo)
	repo.db = db
	return &EventRepo{
		Repo: repo,
		row:  new(pg_models2.Event),
		rows: []*pg_models2.Event{},
	}
}

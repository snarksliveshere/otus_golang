package pg_repository

import (
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/internal/interfaces/repositories/pg_repository/pg_models"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/tools/logger"
)

type DateRepo struct {
	*Repo
	row  *pg_models.Calendar
	rows []*pg_models.Calendar
}

type EventRepo struct {
	*Repo
	row  *pg_models.Event
	rows []*pg_models.Event
}

type Repo struct {
	db     *pg.DB
	logger logger.Logger
}

func GetDateRepo(db *pg.DB) *DateRepo {
	repo := new(Repo)
	repo.db = db
	return &DateRepo{
		Repo: repo,
		row:  new(pg_models.Calendar),
		rows: []*pg_models.Calendar{},
	}
}

func GetEventRepo(db *pg.DB) *EventRepo {
	repo := new(Repo)
	repo.db = db
	return &EventRepo{
		Repo: repo,
		row:  new(pg_models.Event),
		rows: []*pg_models.Event{},
	}
}

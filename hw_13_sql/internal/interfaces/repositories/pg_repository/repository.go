package pg_repository

import (
	"github.com/go-pg/pg"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/interfaces/repositories/pg_repository/pg_models"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/tools/logger"
)

type DateRepo struct {
	*Repo
	row  *pg_models.Calendar
	rows []*pg_models.Calendar
}

type RecordRepo struct {
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

func GetRecordRepo(db *pg.DB) *RecordRepo {
	repo := new(Repo)
	repo.db = db
	return &RecordRepo{
		Repo: repo,
		row:  new(pg_models.Event),
		rows: []*pg_models.Event{},
	}
}

package pg_repository

import (
	"github.com/snarskliveshere/otus_golang/hw_13_sql/config"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/internal/usecases"
	"github.com/snarskliveshere/otus_golang/hw_13_sql/tools/db"
)

type Storage struct {
	Actions *usecases.Actions
}

func CreateStorageInstance(logger usecases.Logger, conf *config.Config) *Storage {
	dbHandler := db.CreatePgConn(conf, logger)
	actions := new(usecases.Actions)
	actions.Logger = logger
	actions.DateRepository = GetDateRepo(dbHandler)
	actions.RecordRepository = GetRecordRepo(dbHandler)
	return &Storage{Actions: actions}
}

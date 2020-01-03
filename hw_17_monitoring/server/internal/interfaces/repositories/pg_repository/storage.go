package pg_repository

import (
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/config"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/internal/pkg/databases/postgres"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/internal/usecases"
)

type Storage struct {
	Actions *usecases.Actions
}

func CreateStorageInstance(logger usecases.Logger, conf *config.AppConfig) *Storage {
	db := postgres.DB{
		Conf: conf,
		Log:  logger,
	}
	dbHandler := db.CreatePgConn()
	actions := new(usecases.Actions)
	actions.Logger = logger
	actions.DateRepository = GetDateRepo(dbHandler)
	actions.EventRepository = GetEventRepo(dbHandler)
	return &Storage{Actions: actions}
}

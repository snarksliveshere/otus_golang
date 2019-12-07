package pg_repository

import (
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/config"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/internal/usecases"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/tools/db"
)

type Storage struct {
	Actions *usecases.Actions
}

func CreateStorageInstance(logger usecases.Logger, conf *config.Config) *Storage {
	dbHandler := db.CreatePgConn(conf, logger)
	actions := new(usecases.Actions)
	actions.Logger = logger
	actions.DateRepository = GetDateRepo(dbHandler)
	actions.EventRepository = GetEventRepo(dbHandler)
	return &Storage{Actions: actions}
}

package pg_repository

import (
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/config"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/internal/usecases"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/pkg/databases/postgres"
)

type Storage struct {
	Actions *usecases.Actions
}

func CreateStorageInstance(logger usecases.Logger, conf *config.Config) *Storage {
	dbHandler := postgres.CreatePgConn(conf, logger)
	actions := new(usecases.Actions)
	actions.Logger = logger
	actions.DateRepository = GetDateRepo(dbHandler)
	actions.EventRepository = GetEventRepo(dbHandler)
	return &Storage{Actions: actions}
}

package postgres

import (
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/config"
	"github.com/snarksliveshere/otus_golang/hw_15_docker/server/internal/usecases"
	"sync"
)

const dbSchema = "calendar"

var dbOnce sync.Once
var dbConn *pg.DB

func CreatePgConn(conf *config.Config, log usecases.Logger) *pg.DB {
	opt := &pg.Options{
		Addr:     conf.DbHost + ":" + conf.DbPort,
		User:     conf.DbUser,
		Password: conf.DbPassword,
		Database: conf.DbName,
	}
	dbOnce.Do(func() {
		dbConn = pg.Connect(opt)
		if _, err := dbConn.Exec("set search_path=?", dbSchema); err != nil {
			log.Infof(err.Error())
		}
	})
	return dbConn
}

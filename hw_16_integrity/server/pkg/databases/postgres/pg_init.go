package postgres

import (
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/server/config"
	"github.com/snarksliveshere/otus_golang/hw_16_integrity/server/internal/usecases"
	"sync"
)

const dbSchema = "calendar"

var dbOnce sync.Once
var dbConn *pg.DB

func CreatePgConn(conf *config.AppConfig, log usecases.Logger) *pg.DB {
	opt := &pg.Options{
		Addr:     conf.DBHost + ":" + conf.DBPort,
		User:     conf.DBUser,
		Password: conf.DBPassword,
		Database: conf.DBName,
	}

	dbOnce.Do(func() {
		dbConn = pg.Connect(opt)
		if _, err := dbConn.Exec("set search_path=?", dbSchema); err != nil {
			log.Infof(err.Error())
		}
	})
	return dbConn
}

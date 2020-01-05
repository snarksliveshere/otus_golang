package model

import (
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/message_office/config"
	"log"
	"sync"
)

const dbSchema = "calendar"

var dbOnce sync.Once
var dbConn *pg.DB

type DB struct {
	Conf *config.AppConfig
}

func (db *DB) CreatePgConn() *pg.DB {
	opt := &pg.Options{
		Addr:     db.Conf.DBHost + ":" + db.Conf.DBPort,
		User:     db.Conf.DBUser,
		Password: db.Conf.DBPassword,
		Database: db.Conf.DBName,
	}

	dbOnce.Do(func() {
		dbConn = pg.Connect(opt)
		if _, err := dbConn.Exec("set search_path=?", dbSchema); err != nil {
			log.Print(err.Error())
		}
	})
	return dbConn
}

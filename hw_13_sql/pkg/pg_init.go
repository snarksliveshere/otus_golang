package pkg

import (
	"github.com/go-pg/pg"
	"sync"
)

const dbSchema = "calendar"

var dbOnce sync.Once
var dbConn *pg.DB

func Db() *pg.DB {

	opt := pg.Options{
		Network:               "",
		Addr:                  "",
		Dialer:                nil,
		OnConnect:             nil,
		User:                  "",
		Password:              "",
		Database:              "",
		ApplicationName:       "",
		TLSConfig:             nil,
		MaxRetries:            0,
		RetryStatementTimeout: false,
		MinRetryBackoff:       0,
		MaxRetryBackoff:       0,
		DialTimeout:           0,
		ReadTimeout:           0,
		WriteTimeout:          0,
		PoolSize:              0,
		MinIdleConns:          0,
		MaxConnAge:            0,
		PoolTimeout:           0,
		IdleTimeout:           0,
		IdleCheckFrequency:    0,
	}
	dbOnce.Do(func() {
		dbConn = pg.Connect(pg_ext.ConnOptsFromDsn(Conf().DbDsn))
		if _, err := dbConn.Exec("set search_path=?", dbSchema); err != nil {
			Log().Panic(err)
		}
		dbConn.AddQueryHook(pg_ext.DbLogger{
			LogFunc: func(q string, p []interface{}) { Log().Debugf("query: %s", q) },
			ErrFunc: func(err error) { Log().Panic(err) },
		})
	})
	return dbConn
}

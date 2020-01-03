package main

import (
	"flag"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	"github.com/kelseyhightower/envconfig"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/config"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/internal/pkg/databases/postgres"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/internal/pkg/logger/logrus"
	"github.com/snarksliveshere/otus_golang/hw_17_monitoring/server/internal/usecases"
	"log"
)

func main() {
	var conf config.AppConfig
	failOnError(envconfig.Process("reg_service", &conf), "failed to init config")
	logg := logrus.CreateLogrusLog(conf.LogLevel)
	logg.Infof("Configs are %#v", conf)

	db := postgres.DB{
		Conf: &conf,
		Log:  logg,
	}.CreatePgConn()
	initMigrationTableIfNeeded(db, logg)
	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		log.Fatal(err.Error())
	}
	if newVersion != oldVersion {
		logg.Infof("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		logg.Infof("version is %d\n", oldVersion)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func initMigrationTableIfNeeded(db *pg.DB, log usecases.Logger) {
	schema, err := getCurrentSchema(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	setSchemaForMigration(schema)
	if err != nil {

	}
	exist, err := isExistMigrationTable(db, schema)
	if err != nil {
		log.Fatal(err.Error())
	}
	if !exist {
		_, _, err = migrations.Run(db, "init")
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func getCurrentSchema(db *pg.DB) (schema string, err error) {
	_, err = db.Query(pg.Scan(&schema), "show search_path")
	return schema, err
}

func setSchemaForMigration(schema string) {
	migrations.SetTableName(schema + "." + config.MigrationTable)
}

func isExistMigrationTable(db *pg.DB, schema string) (exist bool, err error) {
	_, err = db.Query(pg.Scan(&exist), "SELECT EXISTS ("+
		"SELECT 1 "+
		"FROM   information_schema.tables "+
		"WHERE  table_schema = ?0 "+
		"AND    table_name = ?1"+
		");", schema, config.MigrationTable)
	return exist, err
}

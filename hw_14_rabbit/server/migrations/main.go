package main

import (
	"flag"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/config"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/internal/usecases"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/pkg/databases/postgres"
	"github.com/snarksliveshere/otus_golang/hw_14_rabbit/server/pkg/logger/logrus"
	"os"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	conf := config.CreateConfig("./config/config.yaml")
	log := logrus.CreateLogrusLog(conf)
	db := postgres.CreatePgConn(conf, log)
	initMigrationTableIfNeeded(db, log)
	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		log.Fatal(err.Error())
	}
	if newVersion != oldVersion {
		log.Infof("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		log.Infof("version is %d\n", oldVersion)
	}
}

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
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

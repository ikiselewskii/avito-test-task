package database

import (
	"github.com/ikiselewskii/avito-test-task/utils"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
)

var DB *bun.DB

func InitializeDBConnection(dsn string) *bun.DB {
	config, err := pgx.ParseConnectionString(utils.SerializeDSN())
	if err != nil {
		log.Fatal("can`t parse DSN")
	}
	config.PreferSimpleProtocol = true
	sqldb := stdlib.OpenDB(config)
	err = sqldb.Ping()
	if err != nil {
		log.Fatal(err)
	}
	DB = bun.NewDB(sqldb, pgdialect.New())
	DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return DB
}

package database

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var DB *bun.DB
//TODO переписать на pgx-драйвер по совету из статьи avito-tech c медиума
//TODO установка макс количества соединений из конфигов 
func InitializeDBConnection(dsn string) *bun.DB{
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	DB = bun.NewDB(sqldb, pgdialect.New())
	DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return DB
}

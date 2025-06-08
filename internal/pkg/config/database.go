package config

import (
	"github.com/yusologia/go-core/v2/database"
	"gorm.io/gorm"
	"os"
)

var (
	PgSQL *gorm.DB
)

func InitDB() func() {
	var DBClose func()

	PgSQL, DBClose = logiadb.Connect(logiadb.DBConf{
		Driver:    logiadb.POSTGRESQL_DRIVER,
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		Username:  os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		Database:  os.Getenv("DB_DATABASE"),
		ParseTime: true,
	})

	return DBClose
}

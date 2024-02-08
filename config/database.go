package config

import (
	"github.com/yusologia/go-core/config"
	"gorm.io/gorm"
	"os"
)

var (
	PgSQL *gorm.DB
)

func InitDB() {
	PgSQL = config.Connect(config.DBConf{
		Driver:    config.POSTGRESQL_DRIVER,
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		Username:  os.Getenv("DB_USERNAME"),
		Password:  os.Getenv("DB_PASSWORD"),
		Database:  os.Getenv("DB_DATABASE"),
		ParseTime: true,
	})
}

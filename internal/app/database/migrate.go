package database

import (
	"github.com/yusologia/go-core/v2/database"
	"service/internal/app/database/migration"
)

func Migrations() []logiadb.Migration {
	return []logiadb.Migration{
		&migration.Activity_1726651211960757{},
	}
}

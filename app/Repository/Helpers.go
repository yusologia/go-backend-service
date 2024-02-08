package Repository

import (
	"fmt"
	"github.com/yusologia/go-core/helpers/logialog"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"service/config"
	"time"
)

func GetIncrementMonthly(model interface{}) int64 {
	var totalData int64
	config.PgSQL.Unscoped().
		Where("EXTRACT(MONTH FROM \"createdAt\") = ?", time.Now().Month()).
		Model(&model).
		Count(&totalData)

	return totalData + 1
}

func Truncate(db *gorm.DB, tables ...schema.Tabler) {
	if len(tables) > 0 {
		for _, table := range tables {
			err := db.Exec(fmt.Sprintf("truncate table %s restart identity cascade", table.TableName()))
			if err != nil {
				logialog.Error(fmt.Sprintf("Truncate invalid: %v", err))
			}
		}
	}
}

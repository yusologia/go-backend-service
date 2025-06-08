package model

import (
	"github.com/yusologia/go-core/v2/model"
)

type TestingSub struct {
	logiamodel.BaseModel
	TestingId uint   `gorm:"column:testingId;type:bigint;not null"`
	Name      string `gorm:"column:name;type:varchar(250);default:null"`
}

func (TestingSub) TableName() string {
	return "testing_subs"
}

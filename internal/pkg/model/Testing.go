package model

import (
	"github.com/yusologia/go-core/v2/model"
)

type Testing struct {
	logiamodel.BaseModel
	Name string `gorm:"column:name;type:varchar(250);default:null"`

	Subs []TestingSub `gorm:"foreignKey:testingId"`
}

func (Testing) TableName() string {
	return "testing"
}

func (model Testing) SetReference() uint {
	return model.BaseModel.ID
}

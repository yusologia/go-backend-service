package model

import (
	"github.com/yusologia/go-core/v2/model"
)

type Activity struct {
	logiamodel.BaseModel
	Action        string                        `gorm:"column:action;type:varchar(50);not null"`
	SubFeature    string                        `gorm:"column:subFeature;type:varchar(150)"`
	Description   string                        `gorm:"column:description;type:text"`
	ReferenceID   uint                          `gorm:"column:referenceId;type:bigint"`
	ReferenceType string                        `gorm:"column:referenceType;type:varchar(150)"`
	CausedBy      string                        `gorm:"column:causedBy;type:varchar(50)"`
	CausedByName  string                        `gorm:"column:causedByName;type:varchar(150)"`
	Properties    logiamodel.MapInterfaceColumn `gorm:"column:properties;type:json"`
}

func (Activity) TableName() string {
	return "activities"
}

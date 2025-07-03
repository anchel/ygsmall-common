package models_admin

import "gorm.io/gorm"

type ImageDirectory struct {
	gorm.Model

	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

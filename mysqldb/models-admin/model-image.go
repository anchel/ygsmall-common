package models_admin

import "gorm.io/gorm"

type Image struct {
	gorm.Model

	DirectoryID uint   `gorm:"column:directory_id;type:bigint" json:"directory_id"`
	Name        string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Size        string `gorm:"column:size;type:varchar(255)" json:"size"`
	URL         string `gorm:"column:url;type:varchar(2048);not null" json:"url"`
}

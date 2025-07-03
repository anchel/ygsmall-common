package models

import "time"

type CategoryStatus string

const (
	CategoryStatusActive   CategoryStatus = "active"
	CategoryStatusInactive CategoryStatus = "inactive"
)

type Category struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	Name      string         `gorm:"column:name;type:varchar(255);not null;unique_index" json:"name"`
	ParentID  int64          `gorm:"column:parent_id;type:bigint;not null;default:0" json:"parent_id"`
	Level     int32          `gorm:"column:level;type:int;not null;default:0" json:"level"`
	Status    CategoryStatus `gorm:"column:status;type:varchar(255);not null;default:'active'" json:"status"`
	SortOrder int32          `gorm:"column:sort_order;type:int;not null;default:0" json:"sort_order"`
	Thumbnail string         `gorm:"column:thumbnail;type:varchar(2048)" json:"thumbnail"`

	Attributes []CategoryAttribute `gorm:"foreignKey:CategoryID;references:ID" json:"attributes"`
}

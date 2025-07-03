package models

import "time"

type CategoryAttribute struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	CategoryID  int64 `gorm:"column:category_id;type:bigint;not null" json:"category_id"`
	AttributeID int64 `gorm:"column:attribute_id;type:bigint;not null" json:"attribute_id"`
	SortOrder   int32 `gorm:"column:sort_order;type:tinyint;not null;default:0" json:"sort_order"`
}

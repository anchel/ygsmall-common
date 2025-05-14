package models

import "time"

type ProductAttribute struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	ProductID        uint   `gorm:"column:product_id;type:bigint;not null" json:"product_id"`
	AttributeID      uint   `gorm:"column:attribute_id;type:bigint;not null" json:"attribute_id"`
	AttributeValueId string `gorm:"column:attribute_value_id;type:bigint;not null" json:"attribute_value_id"`
	SortOrder        int32  `gorm:"column:sort_order;type:int;not null;default:0" json:"sort_order"`
}

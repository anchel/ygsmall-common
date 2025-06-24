package models

import "time"

type ProductSkuAttribute struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	SkuID int64 `gorm:"column:sku_id;type:bigint;not null" json:"sku_id"`

	AttributeID int64      `gorm:"column:attribute_id;type:bigint;not null" json:"attribute_id"`
	Attribute   *Attribute `gorm:"foreignKey:AttributeID;references:ID" json:"attribute"`

	AttributeValueID int64           `gorm:"column:attribute_value_id;type:bigint;not null" json:"attribute_value_id"`
	AttributeValue   *AttributeValue `gorm:"foreignKey:AttributeValueID;references:ID" json:"attribute_value"`
}

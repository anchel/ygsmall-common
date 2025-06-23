package models

import "time"

type AttributeValue struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	AttributeID int64 `gorm:"column:attribute_id;type:bigint;not null" json:"attribute_id"`

	Value string `gorm:"column:value;type:varchar(255);not null" json:"value"`

	Attribute Attribute `gorm:"foreignKey:AttributeID;references:ID" json:"attribute"`
}

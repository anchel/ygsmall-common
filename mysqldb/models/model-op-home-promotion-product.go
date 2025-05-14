package models

import "time"

type OpHomePromotionProduct struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	ProductID int64  `gorm:"column:product_id;type:bigint;not null;default:0" json:"product_id"`
	SortOrder int32  `gorm:"column:sort_order;type:int;not null;default:0" json:"sort_order"`
	Status    string `gorm:"column:status;type:varchar(255);not null;default:'active'" json:"status"`
}

package models

import "time"

type ShoppingCart struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	UserID    int64 `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	ProductID int64 `gorm:"column:product_id;type:bigint;not null" json:"product_id"`
	SkuID     int64 `gorm:"column:sku_id;type:bigint;not null" json:"sku_id"`
	Quantity  int64 `gorm:"column:quantity;type:int;not null" json:"quantity"`
	Selected  *bool `gorm:"column:selected;type:tinyint(1);default:0" json:"selected"` // 是否选中
}

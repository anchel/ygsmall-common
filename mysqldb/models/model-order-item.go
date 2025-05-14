package models

import "time"

type OrderItem struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	OrderID string `gorm:"column:order_id;type:varchar(128);not null" json:"order_id"`

	ProductID int64  `gorm:"column:product_id;type:bigint;not null" json:"product_id"`
	SkuID     int64  `gorm:"column:sku_id;type:bigint;not null" json:"sku_id"`
	SkuTitle  string `gorm:"column:sku_title;type:varchar(512);not null" json:"sku_title"`
	SkuImage  string `gorm:"column:sku_image;type:varchar(512);not null" json:"sku_image"`
	SkuSpecs  string `gorm:"column:sku_specs;type:varchar(512);not null" json:"sku_specs"`

	Quantity   int64 `gorm:"column:quantity;type:int;not null" json:"quantity"`
	UnitPrice  int64 `gorm:"column:unit_price;type:bigint;not null" json:"unit_price"`
	TotalPrice int64 `gorm:"column:total_price;type:bigint;not null" json:"total_price"`
}

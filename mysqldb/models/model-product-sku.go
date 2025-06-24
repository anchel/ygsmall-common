package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
)

type ProductSkuStatus string

const (
	ProductSkuStatusActive   ProductSkuStatus = "active"
	ProductSkuStatusInactive ProductSkuStatus = "inactive"
)

type ProductSku struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	ProductID int64 `gorm:"column:product_id;type:bigint;not null" json:"product_id"`

	Title    string     `gorm:"column:title;type:varchar(255);not null;unique_index" json:"title"`
	SkuImage string     `gorm:"column:sku_image;type:varchar(2048)" json:"sku_image"`
	Details  SkuDetails `gorm:"column:details;type:json" json:"details"`

	Attributes []ProductSkuAttribute `gorm:"foreignKey:ID;references:SkuID" json:"attributes"`

	Status ProductSkuStatus `gorm:"column:status;type:varchar(255);not null;default:'inactive'" json:"status"`

	PriceOrigin   int64 `gorm:"column:price_origin;type:int;not null" json:"price_origin"`
	PriceSale     int64 `gorm:"column:price_sale;type:int;not null" json:"price_sale"`
	StockQuantity int64 `gorm:"column:stock_quantity;type:int;not null" json:"stock_quantity"`
}

type SkuDetails struct {
	List []SkuDetailsItem `json:"list"`
}

type SkuDetailsItem struct {
	Typ      string `json:"type"` // image
	ImageUrl string `json:"image_url"`
}

// 实现 driver.Valuer 接口（写入 DB）
func (p SkuDetails) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// 实现 sql.Scanner 接口（读取 DB）
func (p *SkuDetails) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		log.Error("SkuDetails Failed to scan JSON from DB", "value", value)
		return fmt.Errorf("SkuDetails Failed to unmarshal JSON from DB")
	}
	return json.Unmarshal(bytes, p)
}

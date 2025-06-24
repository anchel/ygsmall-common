package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
)

type ProductStatus string

const (
	ProductStatusActive   ProductStatus = "active"
	ProductStatusInactive ProductStatus = "inactive"
)

type Product struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	Title        string         `gorm:"column:title;type:varchar(255);not null;unique_index" json:"title"`
	Description  string         `gorm:"column:description;type:varchar(2048);not null" json:"description"`
	PrimaryImage string         `gorm:"column:primary_image;type:varchar(2048)" json:"primary_image"`
	Details      ProductDetails `gorm:"column:details;type:json" json:"details"`

	CategoryID int64     `gorm:"column:category_id;type:bigint;not null" json:"category_id"`
	Category   *Category `gorm:"foreignKey:CategoryID;references:ID" json:"category"`

	Attributes []ProductAttribute `gorm:"foreignKey:ProductID;references:ID" json:"attributes"`
	SkuList    []ProductSku       `gorm:"foreignKey:ProductID;references:ID" json:"sku_list"`

	Status ProductStatus `gorm:"column:status;type:varchar(255);not null;default:'inactive'" json:"status"`

	PriceOrigin   int32 `gorm:"column:price_origin;type:int;not null" json:"price_origin"`
	PriceSale     int32 `gorm:"column:price_sale;type:int;not null" json:"price_sale"`
	StockQuantity int32 `gorm:"column:stock_quantity;type:int;not null" json:"stock_quantity"`
}

type ProductDetails struct {
	List []SkuDetailsItem `json:"list"`
}

type ProductDetailsItem struct {
	Typ      string `json:"type"` // image
	ImageUrl string `json:"image_url"`
}

// 实现 driver.Valuer 接口（写入 DB）
func (p ProductDetails) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// 实现 sql.Scanner 接口（读取 DB）
func (p *ProductDetails) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		log.Error("SkuDetails Failed to scan JSON from DB", "value", value)
		return fmt.Errorf("SkuDetails Failed to unmarshal JSON from DB")
	}
	return json.Unmarshal(bytes, p)
}

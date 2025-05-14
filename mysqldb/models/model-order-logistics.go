package models

import "time"

type OrderLogisticsType string

const (
	OrderLogisticsTypeNormal      OrderLogisticsType = "normal"       // 正常商家发货
	OrderLogisticsTypeReturnGoods OrderLogisticsType = "return_goods" // 退货
)

type OrderLogistics struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	OrderID string `gorm:"column:order_id;type:varchar(128);not null" json:"order_id"`

	Typ                  OrderLogisticsType `gorm:"column:type;type:varchar(32);not null" json:"type"` // normal-正常商家发货 return_goods-退货
	LogisticsCompanyCode string             `gorm:"column:logistics_company_code;type:varchar(32);not null" json:"logistics_company_code"`
	LogisticsCompanyName string             `gorm:"column:logistics_company_name;type:varchar(128);not null" json:"logistics_company_name"`
	LogisticsNo          string             `gorm:"column:logistics_no;type:varchar(64);not null" json:"logistics_no"`
	Remark               string             `gorm:"column:remark;type:varchar(255);not null" json:"remark"` // 备注
}

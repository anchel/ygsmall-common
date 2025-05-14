package models

import "time"

type OrderAddress struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	OrderID string `gorm:"column:order_id;type:varchar(128);not null" json:"order_id"`

	RecipientName string `gorm:"column:recipient_name;type:varchar(255);not null" json:"recipient_name"`
	Phone         string `gorm:"column:phone;type:varchar(255);not null" json:"phone"`
	Country       string `gorm:"column:country;type:varchar(255);not null" json:"country"`
	Province      string `gorm:"column:province;type:varchar(255);not null" json:"province"`
	City          string `gorm:"column:city;type:varchar(255);not null" json:"city"`
	District      string `gorm:"column:district;type:varchar(255);not null" json:"district"`
	DetailAddress string `gorm:"column:detail_address;type:varchar(255);not null" json:"detail_address"`
	PostalCode    string `gorm:"column:postal_code;type:varchar(255);not null" json:"postal_code"`
}

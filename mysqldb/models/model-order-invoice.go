package models

import "time"

type OrderInvoice struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	OrderID string `gorm:"column:order_id;type:varchar(128);not null" json:"order_id"`

	InvoiceType int64  `gorm:"column:invoice_type;type:tinyint;not null" json:"invoice_type"` // 3-增值税电子发票
	Name        string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Taxno       string `gorm:"column:taxno;type:varchar(255);not null" json:"taxno"`
	Phone       string `gorm:"column:phone;type:varchar(255);not null" json:"phone"`
	Email       string `gorm:"column:email;type:varchar(255);not null" json:"email"`
	TitleType   int64  `gorm:"column:title_type;type:tinyint;not null" json:"title_type"`     // 1-个人 2-公司
	ContentType int64  `gorm:"column:content_type;type:tinyint;not null" json:"content_type"` // 1-明细 2-类别
}

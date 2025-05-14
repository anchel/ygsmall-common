package models

import (
	"time"
)

type OrderInfoStatusType string

const (
	OrderInfoStatusTypeActive   OrderInfoStatusType = "active"   // 活动中
	OrderInfoStatusTypeInactive OrderInfoStatusType = "inactive" // 失效
)

type OrderInfo struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	UserID     int64               `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	OrderID    string              `gorm:"column:order_id;type:varchar(128);not null" json:"order_id"`
	ExpireTime time.Time           `gorm:"column:expire_time;type:datetime;not null" json:"expire_time"`
	Status     OrderInfoStatusType `gorm:"column:status;type:varchar(64);not null" json:"status"`
}

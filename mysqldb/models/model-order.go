package models

import (
	"time"
)

type OrderStatus int32

func (v OrderStatus) String() string {
	switch v {
	case OrderStatusUnpaid:
		return "待付款"
	case OrderStatusPaid:
		return "待发货"
	case OrderStatusShipped:
		return "待收货"
	case OrderStatusCompleted:
		return "已完成"
	case OrderStatusAfterSalesService:
		return "退款/售后"
	case OrderStatusCancelled:
		return "已取消"
	default:
		return "未知状态"
	}
}

const (
	OrderStatusUnpaid    OrderStatus = 0 // 待付款
	OrderStatusPaid      OrderStatus = 1 // 待发货 （已付款）
	OrderStatusShipped   OrderStatus = 2 // 待收货 （已发货）
	OrderStatusCompleted OrderStatus = 3 // 已完成

	OrderStatusAfterSalesService OrderStatus = 7 // 退款或售后
	OrderStatusCancelled         OrderStatus = 8 // 已取消，包含：未支付主动取消，超时未支付取消，已支付未发货的主动取消
)

type Order struct {
	ID        string    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	UserID int64 `gorm:"column:user_id;type:bigint;not null" json:"user_id"`

	TotalAmount    int64       `gorm:"column:total_amount;type:bigint;not null" json:"total_amount"`
	ProductAmount  int64       `gorm:"column:product_amount;type:bigint;not null" json:"product_amount"`
	DiscountAmount int64       `gorm:"column:discount_amount;type:bigint;not null" json:"discount_amount"`
	PayAmount      int64       `gorm:"column:pay_amount;type:bigint;not null" json:"pay_amount"`
	Freight        int64       `gorm:"column:freight;type:bigint;not null" json:"freight"`
	Status         OrderStatus `gorm:"column:status;type:int;not null" json:"status"`
	Remark         string      `gorm:"column:remark;type:varchar(512)" json:"remark"`
	ProductDesc    string      `gorm:"column:product_desc;type:varchar(512)" json:"product_desc"`

	PayExpireAt      time.Time  `gorm:"column:pay_expire_at;type:datetime;not null" json:"pay_expire_at"`
	PaidAt           *time.Time `gorm:"column:paid_at;type:datetime" json:"paid_at"`
	ShippedAt        *time.Time `gorm:"column:shipped_at;type:datetime" json:"shipped_at"`
	ConfirmReceiptAt *time.Time `gorm:"column:confirm_receipt_at;type:datetime" json:"confirm_receipt_at"`
}

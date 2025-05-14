package models

import (
	"time"
)

type OrderRefundStatusType string

// PENDING:等待执行 SUCCESS: 退款成功 CLOSED: 退款关闭 PROCESSING: 退款处理中 ABNORMAL: 退款异常
const (
	OrderRefundStatusTypePending    OrderRefundStatusType = "PENDING"    // 等待执行
	OrderRefundStatusTypeProcessing OrderRefundStatusType = "PROCESSING" // 已提交退款
	OrderRefundStatusTypeSuccess    OrderRefundStatusType = "SUCCESS"    // 退款成功
	OrderRefundStatusTypeClosed     OrderRefundStatusType = "CLOSED"     // 已关闭
	OrderRefundStatusTypeAbnormal   OrderRefundStatusType = "ABNORMAL"   // 退款异常
)

type OrderRefundReasonType string

const (
	OrderRefundReasonTypeUserCancel   OrderRefundReasonType = "USER_CANCEL"   // 用户取消
	OrderRefundReasonTypeApplyService OrderRefundReasonType = "APPLY_SERVICE" // 申请售后
)

type OrderRefund struct {
	RefundNO  string    `gorm:"primarykey" json:"refund_no"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	UserID       int64  `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	OrderID      string `gorm:"column:order_id;type:varchar(128);not null" json:"order_id"`
	AfterSalesID string `gorm:"column:after_sales_id;type:varchar(128);not null" json:"after_sales_id"`

	RefundAmount int64                 `gorm:"column:refund_amount;type:bigint;not null" json:"refund_amount"`
	ReasonType   OrderRefundReasonType `gorm:"column:reason_type;type:varchar(64);not null" json:"reason_type"`
	Remark       string                `gorm:"column:remark;type:varchar(64);not null" json:"remark"`

	Status OrderRefundStatusType `gorm:"column:status;type:varchar(64);not null" json:"status"`

	ApplyTime   time.Time  `gorm:"column:apply_time;type:datetime;not null" json:"apply_time"`
	SuccessTime *time.Time `gorm:"column:success_time;type:datetime;not null" json:"success_time"`
}

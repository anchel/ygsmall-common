package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
)

type OrderAftersalesServiceType string

const (
	OrderAftersalesServiceTypeReturnMoney      OrderAftersalesServiceType = "RETURN_MONEY"       // 退款
	OrderAftersalesServiceTypeReturnGoodsMoney OrderAftersalesServiceType = "RETURN_GOODS_MONEY" // 退货退款
)

func (t OrderAftersalesServiceType) String() string {
	switch t {
	case OrderAftersalesServiceTypeReturnMoney:
		return "退款"
	case OrderAftersalesServiceTypeReturnGoodsMoney:
		return "退货退款"
	default:
		return "未知操作"
	}
}

type OrderAftersalesServiceGoodsReceiptStatus string

const (
	OrderAftersalesServiceGoodsReceiptStatusNotReceived OrderAftersalesServiceGoodsReceiptStatus = "NOT_RECEIVED" // 未收货
	OrderAftersalesServiceGoodsReceiptStatusReceived    OrderAftersalesServiceGoodsReceiptStatus = "RECEIVED"     // 已收货
)

func (t OrderAftersalesServiceGoodsReceiptStatus) String() string {
	switch t {
	case OrderAftersalesServiceGoodsReceiptStatusNotReceived:
		return "未收货"
	case OrderAftersalesServiceGoodsReceiptStatusReceived:
		return "已收货"
	default:
		return "未知操作"
	}
}

type OrderAftersalesServiceStatus string

// applied-已申请 canceled-用户取消 approved-审核通过 rejected-商家已拒绝 completed-已完成
const (
	OrderAftersalesServiceStatusApplied   OrderAftersalesServiceStatus = "applied"   // 已申请
	OrderAftersalesServiceStatusCanceled  OrderAftersalesServiceStatus = "canceled"  // 用户取消
	OrderAftersalesServiceStatusApproved  OrderAftersalesServiceStatus = "approved"  // 审核通过
	OrderAftersalesServiceStatusRejected  OrderAftersalesServiceStatus = "rejected"  // 商家已拒绝
	OrderAftersalesServiceStatusCompleted OrderAftersalesServiceStatus = "completed" // 已完成
)

type OrderAftersalesServiceRefundStatus string

// none-无 approved-审核通过 processing-退款中 completed-已完成
const (
	OrderAftersalesServiceRefundStatusNone       OrderAftersalesServiceRefundStatus = "none"       // 无退款
	OrderAftersalesServiceRefundStatusApproved   OrderAftersalesServiceRefundStatus = "approved"   // 审核通过
	OrderAftersalesServiceRefundStatusProcessing OrderAftersalesServiceRefundStatus = "processing" // 退款中
	OrderAftersalesServiceRefundStatusCompleted  OrderAftersalesServiceRefundStatus = "completed"  // 已完成
)

type OrderAftersalesService struct {
	ID        string    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	UserID             int64                                    `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	OrderID            string                                   `gorm:"column:order_id;type:varchar(128);not null" json:"order_id"`
	ServiceType        OrderAftersalesServiceType               `gorm:"column:service_type;type:varchar(32);not null" json:"service_type"`
	GoodsReceiptStatus OrderAftersalesServiceGoodsReceiptStatus `gorm:"column:goods_receipt_status;type:varchar(32);not null" json:"goods_receipt_status"`
	ReasonType         string                                   `gorm:"column:reason_type;type:varchar(64);not null" json:"reason_type"`
	ReasonDesc         string                                   `gorm:"column:reason_desc;type:varchar(512);not null" json:"reason_desc"`
	RefundAmount       int64                                    `gorm:"column:refund_amount;type:bigint;not null" json:"refund_amount"`
	Remark             string                                   `gorm:"column:remark;type:varchar(512);not null" json:"remark"`
	Evidence           *EvidenceDetails                         `gorm:"column:evidence;type:varchar(512);not null" json:"evidence"`

	ApplyTime      time.Time  `gorm:"column:apply_time;type:datetime" json:"apply_time"`
	ApprovedTime   *time.Time `gorm:"column:approved_time;type:datetime" json:"approved_time"`
	RejectedTime   *time.Time `gorm:"column:rejected_time;type:datetime" json:"rejected_time"`
	RejectedRemark *string    `gorm:"column:rejected_remark;type:varchar(512)" json:"rejected_remark"`

	LogisticsID *int64 `gorm:"column:logistics_id;type:bigint" json:"logistics_id"`

	Status       OrderAftersalesServiceStatus       `gorm:"column:status;type:varchar(32);not null" json:"status"`
	RefundStatus OrderAftersalesServiceRefundStatus `gorm:"column:refund_status;type:varchar(32);not null" json:"refund_status"`
}

type EvidenceDetailItem struct {
	Typ string `json:"type"`
	Url string `json:"url"`
}

type EvidenceDetails struct {
	List []EvidenceDetailItem `json:"list"`
}

// 实现 driver.Valuer 接口（写入 DB）
func (p EvidenceDetails) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// 实现 sql.Scanner 接口（读取 DB）
func (p *EvidenceDetails) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		log.Error("EvidenceDetails Failed to scan JSON from DB", "value", value)
		return fmt.Errorf("EvidenceDetails Failed to unmarshal JSON from DB")
	}
	return json.Unmarshal(bytes, p)
}

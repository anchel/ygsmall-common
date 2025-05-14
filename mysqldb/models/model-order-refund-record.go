package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/anchel/ygsmall-common/types"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
)

type OrderRefundRecordStatusType string

// SUCCESS: 退款成功 CLOSED: 退款关闭 PROCESSING: 退款处理中 ABNORMAL: 退款异常
const (
	OrderRefundRecordStatusTypeProcessing OrderRefundRecordStatusType = "PROCESSING" // 已提交退款
	OrderRefundRecordStatusTypeSuccess    OrderRefundRecordStatusType = "SUCCESS"    // 退款成功
	OrderRefundRecordStatusTypeClosed     OrderRefundRecordStatusType = "CLOSED"     // 已关闭
	OrderRefundRecordStatusTypeAbnormal   OrderRefundRecordStatusType = "ABNORMAL"   // 退款异常
)

type OrderRefundRecord struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	UserID  int64  `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	OrderID string `gorm:"column:order_id;type:varchar(128);not null" json:"order_id"`

	RefundNO         string                `gorm:"column:refund_no;type:varchar(128);not null" json:"refund_no"`
	RefundPlatform   types.PayPlatformType `gorm:"column:refund_platform;type:varchar(64);not null" json:"refund_platform"`
	TransactionID    *string               `gorm:"column:transaction_id;type:varchar(128)" json:"transaction_id"`
	PlatformRefundID string                `gorm:"column:platform_refund_id;type:varchar(128)" json:"platform_refund_id"`

	AmountTotal    int64  `gorm:"column:amount_total;type:bigint;not null" json:"amount_total"`
	AmountCurrency string `gorm:"column:amount_currency;type:varchar(64);not null" json:"amount_currency"`
	AmountRefund   int64  `gorm:"column:amount_refund;type:bigint;not null" json:"amount_refund"`
	Reason         string `gorm:"column:reason;type:varchar(64);not null" json:"reason"`

	Status OrderRefundRecordStatusType `gorm:"column:status;type:varchar(64);not null" json:"status"`

	Channel             *string    `gorm:"column:channel;type:varchar(64);not null" json:"channel"`
	UserReceivedAccount *string    `gorm:"column:user_received_account;type:varchar(64);not null" json:"user_received_account"`
	SuccessTime         *time.Time `gorm:"column:success_time;type:datetime" json:"success_time"`
	CreateTime          *time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`

	FundsAccount           string `gorm:"column:funds_account;type:varchar(64);not null" json:"funds_account"`
	AmountPayerTotal       *int64 `gorm:"column:amount_payer_total;type:bigint;not null" json:"amount_payer_total"`
	AmountPayerRefund      *int64 `gorm:"column:amount_payer_refund;type:bigint;not null" json:"amount_payer_refund"`
	AmountSettlementTotal  *int64 `gorm:"column:amount_settlement_total;type:bigint;not null" json:"amount_settlement_total"`
	AmountSettlementRefund *int64 `gorm:"column:amount_settlement_refund;type:bigint;not null" json:"amount_settlement_refund"`
	AmountDiscountRefund   *int64 `gorm:"column:amount_discount_refund;type:bigint;not null" json:"amount_discount_refund"`

	AmountRefundFee int64 `gorm:"column:amount_refund_fee;type:bigint;not null" json:"amount_refund_fee"`

	GoodsDetail     *OrderRefundRecordGoodsDetail     `gorm:"column:goods_detail;type:varchar(64);not null" json:"goods_detail"`
	PromotionDetail *OrderRefundRecordPromotionDetail `gorm:"column:promotion_detail;type:varchar(64);not null" json:"promotion_detail"`
}

type OrderRefundRecordGoodsDetailItem struct {
	refunddomestic.GoodsDetail
}

type OrderRefundRecordGoodsDetail struct {
	Details []refunddomestic.GoodsDetail `json:"details"`
}

// 实现 driver.Valuer
func (p OrderRefundRecordGoodsDetail) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// 实现 sql.Scanner
func (p *OrderRefundRecordGoodsDetail) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return json.Unmarshal([]byte(value.(string)), p)
	}
	return json.Unmarshal(bytes, p)
}

type OrderRefundRecordPromotionDetailItem struct {
	refunddomestic.Promotion
}
type OrderRefundRecordPromotionDetail struct {
	Details []refunddomestic.Promotion `json:"details"`
}

// 实现 driver.Valuer
func (p OrderRefundRecordPromotionDetail) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// 实现 sql.Scanner
func (p *OrderRefundRecordPromotionDetail) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return json.Unmarshal([]byte(value.(string)), p)
	}
	return json.Unmarshal(bytes, p)
}

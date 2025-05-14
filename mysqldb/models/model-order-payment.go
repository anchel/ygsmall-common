package models

import (
	"time"

	"github.com/anchel/ygsmall-common/types"
)

type TradeTypeType string

const (
	TradeTypeTypeJSAPI  TradeTypeType = "JSAPI"  // JSAPI支付
	TradeTypeTypeNATIVE TradeTypeType = "NATIVE" // Native支付
	TradeTypeTypeAPP    TradeTypeType = "APP"    // APP支付
	TradeTypeTypeMWEB   TradeTypeType = "MWEB"   // H5支付
)

type TradeStateType string

const (
	TradeStateTypeSuccess TradeStateType = "SUCCESS" // 支付成功
	TradeStateTypeRefund  TradeStateType = "REFUND"  // 转入退款
	TradeStateTypeNotPay  TradeStateType = "NOTPAY"  // 未支付
	TradeStateTypeClosed  TradeStateType = "CLOSED"  // 已关闭
)

type OrderPayment struct {
	TradeNO   string    `gorm:"primarykey;varchar(128);not null" json:"trade_no"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	UserID  int64  `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	OrderID string `gorm:"column:order_id;type:varchar(128);not null" json:"order_id"`

	Total       int64  `gorm:"column:total;type:bigint;not null" json:"total"`
	Currency    string `gorm:"column:currency;type:varchar(64);not null" json:"currency"`
	Description string `gorm:"column:description;type:varchar(255);not null" json:"description"`
	Attach      string `gorm:"column:attach;type:varchar(255);not null" json:"attach"`

	// 下面这些字段，根据实际支付记录来填充
	PayPlatform   *types.PayPlatformType `gorm:"column:pay_platform;type:varchar(64)" json:"pay_platform"`
	TransactionID *string                `gorm:"column:transaction_id;type:varchar(128)" json:"transaction_id"`
	PayerTotal    *int64                 `gorm:"column:payer_total;type:bigint" json:"payer_total"`
	PayerCurrency *string                `gorm:"column:payer_currency;type:varchar(64)" json:"payer_currency"`

	TradeState  TradeStateType `gorm:"column:trade_state;type:varchar(64);not null" json:"trade_state"` // default "NOTPAY"
	SuccessTime *time.Time     `gorm:"column:success_time;type:datetime" json:"success_time"`
}

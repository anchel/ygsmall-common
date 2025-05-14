package models

import (
	"time"

	"github.com/anchel/ygsmall-common/types"
)

type OrderPaymentRecord struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	UserID  int64  `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	OrderID string `gorm:"column:order_id;type:varchar(128);not null" json:"order_id"`
	TradeNO string `gorm:"column:trade_no;type:varchar(128);not null" json:"trade_no"`

	PayPlatform   types.PayPlatformType `gorm:"column:pay_platform;type:varchar(64);not null" json:"pay_platform"`
	PrepayID      string                `gorm:"column:prepay_id;type:varchar(128);not null" json:"prepay_id"`
	TransactionID string                `gorm:"column:transaction_id;type:varchar(128);not null" json:"transaction_id"`

	Total         int64   `gorm:"column:total;type:bigint;not null" json:"total"`
	Currency      string  `gorm:"column:currency;type:varchar(64);not null" json:"currency"`
	PayerTotal    *int64  `gorm:"column:payer_total;type:bigint;not null" json:"payer_total"`
	PayerCurrency *string `gorm:"column:payer_currency;type:varchar(64);not null" json:"payer_currency"`

	AppID       string    `gorm:"column:appid;type:varchar(128);not null" json:"appid"`
	OpenID      string    `gorm:"column:openid;type:varchar(128);not null" json:"openid"`
	MchID       string    `gorm:"column:mchid;type:varchar(128);not null" json:"mchid"`
	TimeExpire  time.Time `gorm:"column:time_expire;type:datetime;not null" json:"time_expire"`
	Description string    `gorm:"column:description;type:varchar(255);not null" json:"description"`
	Attach      string    `gorm:"column:attach;type:varchar(255);not null" json:"attach"`

	TradeType      TradeTypeType  `gorm:"column:trade_type;type:varchar(64);not null" json:"trade_type"`
	TradeState     TradeStateType `gorm:"column:trade_state;type:varchar(64);not null" json:"trade_state"`
	TradeStateDesc string         `gorm:"column:trade_state_desc;type:varchar(255);not null" json:"trade_state_desc"`
	BankType       string         `gorm:"column:bank_type;type:varchar(64);not null" json:"bank_type"`
	SuccessTime    *time.Time     `gorm:"column:success_time;type:datetime;not null" json:"success_time"`
}

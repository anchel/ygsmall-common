package types

import "encoding/json"

type ContextKey string
type ContextValue string

type SessionInfo struct {
	OpenID string `json:"openid" bson:"openid"`

	UserID int64 `json:"user_id" bson:"user_id"`
}

func (s *SessionInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SessionInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

type LoginAuthType string

const (
	LoginAuthTypePhone      LoginAuthType = "phone"
	LoginAuthTypeEmail      LoginAuthType = "email"
	LoginAuthTypeWechatFwh  LoginAuthType = "wechat_fuwuhao"
	LoginAuthTypeWechatMini LoginAuthType = "wechat_miniprogram"
)

type PayPlatformType string

const (
	PayPlatformTypeWechat PayPlatformType = "wechat"
	PayPlatformTypeAlipay PayPlatformType = "alipay"
)

type WechatpayTransactionInfo struct {
	Currency      *string `json:"currency,omitempty"`
	PayerCurrency *string `json:"payer_currency,omitempty"`
	PayerTotal    *int64  `json:"payer_total,omitempty"`
	Total         *int64  `json:"total,omitempty"`

	Appid      *string `json:"appid,omitempty"`
	Attach     *string `json:"attach,omitempty"`
	BankType   *string `json:"bank_type,omitempty"`
	Mchid      *string `json:"mchid,omitempty"`
	OutTradeNo *string `json:"out_trade_no,omitempty"`

	Openid *string `json:"openid,omitempty"`

	SuccessTime    *string `json:"success_time,omitempty"`
	TradeState     *string `json:"trade_state,omitempty"`
	TradeStateDesc *string `json:"trade_state_desc,omitempty"`
	TradeType      *string `json:"trade_type,omitempty"`
	TransactionId  *string `json:"transaction_id,omitempty"`
}

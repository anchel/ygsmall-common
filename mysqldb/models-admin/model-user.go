package models_admin

import "time"

type User struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	AccountName string `gorm:"column:account_name;type:varchar(255);not null" json:"account_name"`
	Email       string `gorm:"column:email;type:varchar(255);not null" json:"email"`
	Nickname    string `gorm:"column:nickname;type:varchar(128);not null" json:"nickname"`
	Avatar      string `gorm:"column:avatar;type:varchar(1024);not null" json:"avatar"`
	Gender      string `gorm:"column:gender;type:varchar(255);not null" json:"gender"`
}

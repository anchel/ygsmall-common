package models_admin

import (
	"time"

	"github.com/anchel/ygsmall-common/types"
)

type UserAuth struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	UserID     int64               `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	AuthType   types.LoginAuthType `gorm:"column:auth_type;type:varchar(255);not null" json:"auth_type"`
	Identifier string              `gorm:"column:identifier;type:varchar(255);not null" json:"identifier"`
	Password   string              `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Salt       string              `gorm:"column:salt;type:varchar(255);not null" json:"salt"`
}

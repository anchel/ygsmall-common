package models

import (
	"time"
)

type UserSearchHistory struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	UserID     int64     `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	SearchWord string    `gorm:"column:search_word;type:varchar(255);not null" json:"search_word"`
	SearchTime time.Time `gorm:"column:search_time;type:datetime;not null" json:"search_time"`
}

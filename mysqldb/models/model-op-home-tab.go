package models

import "time"

type OpHomeTab struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	Title     string `gorm:"column:title;type:varchar(255);not null;unique_index" json:"title"`
	SortOrder int32  `gorm:"column:sort_order;type:int;not null;default:0" json:"sort_order"`
	Status    string `gorm:"column:status;type:varchar(255);not null;default:'active'" json:"status"`
}

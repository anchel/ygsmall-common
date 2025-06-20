package models

import "time"

type Attribute struct {
	ID        int64     `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
}

package model

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID         uuid.UUID  `gorm:"unique_index;not null"`
	CreateTime int64      `gorm:"column:create_time"`
	UpdateTime int64      `gorm:"column:update_time"`
	DeleteTime *time.Time `json:"-" gorm:"type:timestamp;column:delete_time"`
}

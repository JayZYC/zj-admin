package model

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID         uuid.UUID  `gorm:"unique_index;not null"`
	CreateTime *time.Time `gorm:"column:create_time"`
	UpdateTime *time.Time `gorm:"column:update_time"`
	DeleteTime *time.Time `json:"-" gorm:"type:timestamp;column:delete_time"`
}

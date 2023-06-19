package model

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID         uuid.UUID `gorm:"unique_index;not null"`
	CreatTime  int64
	UpdateTime int64
	DeleteTime *time.Time `json:"-" gorm:"type:timestamp;column:deleted"`
}

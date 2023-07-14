package model

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Model struct {
	ID         uuid.UUID      `json:"id" gorm:"unique_index;not null" form:"id" binding:"required"`
	CreateTime *time.Time     `json:"create_time" gorm:"column:create_time;autoCreateTime" form:"create_time"`
	UpdateTime *time.Time     `json:"update_time" gorm:"column:update_time;autoUpdateTime" form:"update_time"`
	DeleteTime gorm.DeletedAt `json:"-" gorm:"index;column:delete_time"` // gorm.DeletedAt gorm框架一个特殊字段，用于支持软删除（Soft Delete）功能
}

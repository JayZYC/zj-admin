package model

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Model struct {
	ID         uuid.UUID      `json:"id" gorm:"primarykey" form:"id"`
	CreateTime *time.Time     `json:"create_time" gorm:"autoCreateTime" form:"create_time"`
	UpdateTime *time.Time     `json:"update_time" gorm:"autoUpdateTime" form:"update_time"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"column:deletedAt"` // gorm.DeletedAt gorm框架一个特殊字段，用于支持软删除（Soft Delete）功能
}

package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type (
	Role struct {
		Model
		Name     string        `json:"name" gorm:"not null;type:varchar(200)"`
		Access   pq.Int64Array `json:"access" gorm:"type:integer[]"` // 权限ID集合
		ParentID uuid.UUID     `json:"parent_id" gorm:"column:parent_id"`
	}

	Roles []Role

	// 更新角色
	UpdateRole struct {
		ID         uuid.UUID     `json:"id" binding:"required"`
		UpdateTime *time.Time    `json:"update_time" gorm:"autoUpdateTime"`
		Name       string        `json:"name"`
		Access     pq.Int64Array `json:"access" gorm:"type:integer[]"` // 权限ID集合
	}
)

func (r *UpdateRole) TableName() string {
	return "role"
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New()
	return
}

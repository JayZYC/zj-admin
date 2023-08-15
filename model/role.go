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

	RolePerm struct {
		ID         uuid.UUID      `json:"id"`
		ParentID   uuid.UUID      `json:"parent_id" gorm:"column:parent_id"`
		Name       string         `json:"name"`
		Type       int            `json:"type"`
		Path       string         `json:"path"`
		Component  string         `json:"component"`
		Perm       string         `json:"perm"`
		Visible    int            `json:"visible"`
		Sort       int            `json:"sort"`
		Icon       string         `json:"icon"`
		Redirec    string         `json:"redirect"`
		CreateTime *time.Time     `json:"create_time"`
		UpdateTime *time.Time     `json:"update_time"`
		DeletedAt  gorm.DeletedAt `json:"deletedAt"`
	}

	// 更新角色
	UpdateRole struct {
		ID         uuid.UUID     `json:"id" binding:"required"`
		UpdateTime *time.Time    `json:"update_time"`
		Name       string        `json:"name"`
		Access     pq.Int64Array `json:"access" gorm:"type:integer[]"` // 权限ID集合
	}
)

package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type (
	Role struct {
		Model
		Name           string        `gorm:"not null;type:varchar(200)"`
		OrganizationID uuid.UUID     `json:"OrganizationID,string"` // 组织ID
		Access         pq.Int64Array `gorm:"type:integer[]"`        // 服务类型
	}

	Roles []Role
)

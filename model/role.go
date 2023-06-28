package model

import (
	"github.com/lib/pq"
)

type (
	Role struct {
		Model
		Name   string        `gorm:"not null;type:varchar(200)"`
		Access pq.Int64Array `gorm:"type:integer[]"` // 服务类型
	}

	Roles []Role
)

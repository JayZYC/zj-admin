package model

import (
	"github.com/google/uuid"
)

type (
	User struct {
		Model
		UserName       string    `gorm:"unique_index;not null;type:varchar(100)"` // 设置用户名和手机号都不允许重复
		Password       []byte    `json:"password,omitempty" gorm:"-;type:"`
		Email          string    `gorm:"type:varchar(300)"`
		Phone          string    `json:"phone,string" gorm:"unique_index"`
		Status         int8      `json:"status,string"`                           // -1：停用  0:正常
		OrganizationID uuid.UUID `json:"organization_id,string"  gorm:"not_null"` // 组织ID
		RoleID         uuid.UUID
		Role           Role
	}

	Users []User
)

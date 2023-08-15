package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	Organization struct {
		Model
		Name     string    `gorm:"type:varchar(800);unique_index;not null;"`
		Province string    `gorm:"type:varchar(200)"`
		City     string    `gorm:"type:varchar(200)"`
		District string    `gorm:"type:varchar(200)"`
		Address  string    `gorm:"type:varchar(800)"`
		ParentID uuid.UUID `json:"parent_id,string"` // 上级组织
	}

	Organizations []*Organization

	OrganizationInfo struct {
		Name        string
		ID          uuid.UUID `gorm:"unique_index;not null"`
		ParentID    uuid.UUID `json:"parent_id,string"` // 上级组织
		TotalDevice int       `gorm:"column:total_device"`
	}

	OrganizationList []*OrganizationInfo

	UpdateOrg struct {
		ID         uuid.UUID  `json:"id" binding:"required"`
		UpdateTime *time.Time `json:"update_time"`
		Name       string     `gorm:"type:varchar(800);unique_index;not null;"`
		Province   string     `gorm:"type:varchar(200)"`
		City       string     `gorm:"type:varchar(200)"`
		District   string     `gorm:"type:varchar(200)"`
		Address    string     `gorm:"type:varchar(800)"`
		ParentID   uuid.UUID  `json:"parent_id,string"` // 上级组织
	}
)

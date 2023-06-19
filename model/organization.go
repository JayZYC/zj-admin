package model

import "github.com/google/uuid"

type (
	Organization struct {
		Model
		Name             string     `gorm:"type:varchar(800);unique_index;not null;"`
		Status           int8       `json:"Status,string"`
		OrganizationType uint8      `json:"OrganizationType,string"`
		Scale            uint8      `json:"Scale,string"`
		Province         string     `gorm:"type:varchar(200)"`
		City             string     `gorm:"type:varchar(200)"`
		District         string     `gorm:"type:varchar(200)"`
		Address          string     `gorm:"type:varchar(800)"`
		ParentID         uuid.UUID  `json:"ParentID,string"` // 上级组织
		Admin            *uuid.UUID `json:"Admin,string"`
		BuildArea        float64    //占地面积
		PersonNum        int        // 单位人数
		Longitude        float64    `json:"Longitude,string"` // 经度
		Latitude         float64    `json:"Latitude,string"`  // 纬度
	}

	Organizations []*Organization

	OrganizationInfo struct {
		Name        string
		ID          uuid.UUID `gorm:"unique_index;not null"`
		ParentID    uuid.UUID `json:"ParentID,string"` // 上级组织
		TotalDevice int       `gorm:"column:total_device"`
	}

	OrganizationList []*OrganizationInfo
)

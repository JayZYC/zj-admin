package model

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID         int64          `json:"id" gorm:"autoIncrement" `
	CreateTime *time.Time     `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime *time.Time     `json:"update_time" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"column:deletedAt"`
	ParentID   int64          `json:"parent_id"`
	Name       string
	Type       int16
	Path       string
	Component  string
	Visible    bool
	Sort       int16
	Icon       string
	Redirect   string
}

func (m *Menu) TableName() string {
	return "role_perm"
}

type MenuRoute struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Path      string `json:"path"`
	*MenuMeta `json:"meta"`
	Children  []MenuRoute `json:"children"`
}

type MenuMeta struct {
	Color  string `json:"color"`
	Icon   string `json:"icon"`
	Title  string `json:"title"`
	IsHide bool   `json:"isHide"`
}

type Menus []Menu

func (m *Menus) TableName() string {
	return "role_perm"
}

type UpdateMenu struct {
	ID         int64      `json:"id" binding:"required"`
	UpdateTime *time.Time `json:"update_time" gorm:"autoUpdateTime"`
	ParentID   int64      `json:"parent_id"`
	Name       string
	Type       int16
	Path       string
	Component  string
	Visible    bool
	Sort       int16
	Icon       string
	Redirect   string
}

func (m *UpdateMenu) TableName() string {
	return "role_perm"
}

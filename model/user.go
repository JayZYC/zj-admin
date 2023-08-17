package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	User struct {
		Model
		UserName       string    `json:"username" gorm:"column:username" form:"username"` // 设置用户名和手机号都不允许重复
		Password       string    `json:"password,omitempty" gorm:"column:password" form:"password"`
		NickName       string    `json:"nick_name" form:"nick_name"`
		Avatar         []byte    `json:"avatar" form:"avatar"`
		Email          string    `gorm:"type:varchar(300)"`
		Phone          string    `json:"phone" gorm:"unique_index"`
		Status         int       `json:"status"`                           //2：停用  1:正常
		OrganizationID uuid.UUID `json:"organization_id"  gorm:"not_null"` // 组织ID
		RoleID         uuid.UUID `json:"role_id" gorm:"column:role_id" form:"role_id"`
	}

	Users []User

	UserInfo struct {
		Model
		UserName string    `json:"username" gorm:"column:username" form:"username"`
		Phone    string    `json:"phone"`
		Email    string    `json:"email"`
		Avatar   []byte    `json:"avatar" form:"avatar"`
		NickName string    `json:"nick_name" form:"nick_name"`
		Status   int       `json:"status"` //2：停用  1:正常
		RoleID   uuid.UUID `json:"role_id" gorm:"column:role_id" form:"role_id"`
		RoleName string    `json:"role_name" form:"role_name"`
	}

	UserInfos []UserInfo

	UpdateUser struct {
		ID             uuid.UUID  `json:"id" gorm:"primarykey" binding:"required"`
		UpdateTime     *time.Time `json:"update_time" gorm:"autoUpdateTime"`
		Password       string     `json:"password,omitempty" gorm:"column:password" form:"password"`
		NickName       string     `json:"nick_name" form:"nick_name"`
		Avatar         []byte     `json:"avatar" form:"avatar"`
		Email          string     `gorm:"type:varchar(300)"`
		Status         int        `json:"status"`                           //2：停用  1:正常
		OrganizationID uuid.UUID  `json:"organization_id"  gorm:"not_null"` // 组织ID
		RoleID         uuid.UUID  `json:"role_id" gorm:"column:role_id" form:"role_id"`
	}
)

func (u *Users) TableName() string {
	return "user"
}

func (u *UpdateUser) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}

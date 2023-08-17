package db

import (
	"zj-admin/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindMenu(id uuid.UUID) (model.Menus, error) {
	var menus model.Menus
	sub := db.Model(&model.Role{}).Select("unnest(access)").Where("id = ?", id)
	err := db.Model(&model.Menus{}).Where("id in (?) and type != ?", sub, 4).Find(&menus).Error
	return menus, err
}

func AddMenu(m model.Menu) error {
	return db.Session(&gorm.Session{SkipHooks: true}).Create(&m).Error
}

func UpdateMenu(m model.UpdateMenu) error {
	return db.Updates(&m).Error
}

func DeleteMenu(id int) error {
	return db.Delete(&model.Menu{}, id).Error
}

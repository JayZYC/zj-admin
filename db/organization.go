package db

import (
	"zj-admin/model"

	"github.com/google/uuid"
)

func FindOrgListByOrgID(id uuid.UUID) (model.Organizations, error) {
	orgs := make(model.Organizations, 0)
	err := db.Raw("SELECT * FROM query_child_organizations(?)", id).Scan(&orgs).Error
	return orgs, err
}

func AddOrg(org model.Organization) error {
	return db.Create(&org).Error
}

func UpdateOrg(org model.UpdateOrg) error {
	return db.Model(&model.Organization{}).Update(&org).Error
}

func DeleteOrg(id uuid.UUID) error {
	return db.Where("id = ?", id).Delete(&model.Organization{}).Error
}

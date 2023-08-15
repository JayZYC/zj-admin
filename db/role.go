package db

import (
	"zj-admin/model"
	"zj-admin/util/pagination"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// 根据role_id查询对应的权限按钮
func FindPermByRoleID(id uuid.UUID) ([]string, error) {
	var perm struct {
		Perm pq.StringArray
	}
	err := db.Raw("SELECT array_agg(perm) as perm FROM role_perm WHERE id = ANY (SELECT unnest(access) FROM role WHERE id = ?) AND perm != ''", id).Scan(&perm).Error
	return perm.Perm, err
}

// 根据role_id查询角色列表
func FindRoleListByRoleID(id uuid.UUID, page pagination.PageRequest) (pagination.PageResponse, error) {
	roles := make(model.Roles, 0)
	var total struct {
		Count int64
	}
	err := db.Raw("SELECT * FROM query_child_roles(?)", id).Offset(page.GetOffset()).Limit(page.PageSize).Scan(&roles).Error
	if err != nil {
		return pagination.PageResponse{Total: 0, List: roles}, err
	}
	err = db.Raw("SELECT count(1) count FROM query_child_roles(?)", id).Scan(&total).Error
	return pagination.PageResponse{Total: total.Count, List: roles}, err
}

func AddRole(role model.Role) error {
	return db.Create(&role).Error
}

func UpdateRole(role model.UpdateRole) error {
	return db.Model(&model.Role{}).Update(&role).Error
}

func DeleteRole(id uuid.UUID) error {
	return db.Where("id = ?", id).Delete(&model.Role{}).Error
}

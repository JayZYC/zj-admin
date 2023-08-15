package db

import (
	"zj-admin/model"
	"zj-admin/util"
	"zj-admin/util/jwt"
	"zj-admin/util/pagination"

	"github.com/google/uuid"
)

func FindUserByNameAndPWD(username, password string) (*model.User, error) {
	var user = new(model.User)
	err := db.Debug().
		Where("username = ? AND password = ?", username, password).
		First(user).Error
	return user, err
}

// 根据登录用户分页查询所有可见用户
func FindUserList(user jwt.CustomClaims, page pagination.PageRequest) (pagination.PageResponse, error) {
	users := make(model.UserInfos, 0)
	var total int64
	var err error

	if util.IsAdmin(user.RoleID) {
		err = db.Table("\"user\" t1").
			Select("t1.id, t1.username, t1.phone, t1.email, t1.avatar, t1.nick_name, t1.role_id, t1.status, t1.create_time, t2.name role_name").
			Joins("LEFT JOIN role t2 ON t1.role_id = t2.id").
			Order("create_time").
			Offset(page.GetOffset()).
			Limit(page.PageSize).
			Scan(&users).
			Offset(-1).
			Limit(-1).
			Count(&total).Error
	} else {
		err = db.Raw("SELECT t1.id, t1.username, t1.phone, t1.email, t1.avatar, t1.nick_name, t1.role_id, t1.status, t1.create_time, t2.name role_name FROM \"user\" t1 LEFT JOIN role t2 ON t1.role_id = t2.id WHERE t1.organization_id = ANY(SELECT query_child_organizations_id(?)) AND t1.\"deletedAt\" IS NULL", user.OrganizationID).
			Order("create_time").
			Offset(page.GetOffset()).
			Limit(page.PageSize).
			Scan(&users).Error
		if err != nil {
			return pagination.PageResponse{Total: total, List: users}, err
		}
		var count struct {
			Count int64
		}
		err = db.Raw("SELECT count(1) count FROM \"user\"  WHERE organization_id = ANY(SELECT query_child_organizations_id(?)) AND \"deletedAt\" IS NULL", user.OrganizationID).Scan(&count).Error
		total = count.Count
	}

	return pagination.PageResponse{Total: total, List: users}, err
}

func AddUser(user model.User) error {
	return db.Create(&user).Error
}

// 更新用户
func UpdateUser(user model.User) error {
	return db.Model(&model.User{}).Omit("create_time", "password", "username", "phone").Update(&user).Error
}

func DeleteUser(id uuid.UUID) error {
	return db.Where("id = ?", id).Delete(&model.User{}).Error
}

// 查询账号是否存在
func UserIsExist(username string) bool {
	var num int64
	db.Model(&model.User{}).Where("username = ?", username).Count(&num)
	return num == 1
}

// 查询手机号是否存在
func PhoneExist(phone string) bool {
	var num int64
	db.Model(&model.User{}).Where("phone = ?", phone).Count(&num)
	return num == 1
}

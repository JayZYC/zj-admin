package db

import "zj-admin/model"

func FindUserByNameAndPWD(username, password string) (*model.User, error) {
	var user = new(model.User)
	err := db.Debug().Where("username = ? AND password = ?", username, password).First(user).Error
	return user, err
}

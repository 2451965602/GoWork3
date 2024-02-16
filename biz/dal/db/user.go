package db

import (
	"context"
	"errors"
	"work3/pkg/constants"
	"work3/pkg/utils/pwd"
)

func CreateUser(ctx context.Context, username, password string) (*User, error) {

	if DB == nil {
		return nil, errors.New("DB object is nil")
	}

	var userResp *User

	err := DB.
		WithContext(ctx).
		Table(constants.UserTable).
		Where("username = ?", username).
		First(&userResp).
		Error

	if err == nil {
		return nil, errors.New("用户名已存在")
	}

	pwd, err := utils.PasswordHash(password)

	if err != nil {
		return nil, err
	}

	userResp = &User{
		Username: username,
		Password: pwd,
	}

	err = DB.
		WithContext(ctx).
		Table(constants.UserTable).
		Create(&userResp).
		Error

	if err != nil {
		return nil, err
	}

	userResp.Password = password

	return userResp, nil
}
func LoginCheck(ctx context.Context, username, password string) (*User, error) {

	var userResp *User

	err := DB.
		WithContext(ctx).
		Table(constants.UserTable).
		Where("username = ?", username).
		First(&userResp).
		Error

	if err != nil {
		return nil, errors.New("账号错误")
	}

	if !utils.PasswordVerify(password, userResp.Password) {
		return nil, errors.New("密码错误")
	}

	//if userResp.Password != password {
	//	return nil, errors.New("密码错误")
	//}

	return userResp, nil
}

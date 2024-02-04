package db

import (
	"context"
	"errors"
	"work3/pkg/constants"
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

	userResp = &User{
		Username: username,
		Password: password,
	}

	err = DB.
		WithContext(ctx).
		Table(constants.UserTable).
		Create(&userResp).
		Error

	if err != nil {
		return nil, err
	}

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

	if userResp.Password != password {
		return nil, errors.New("密码错误")
	}

	return userResp, nil
}

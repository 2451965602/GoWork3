package pack

import (
	"strconv"
	"work3/biz/dal/db"
	"work3/biz/model/model"
)

func User(data *db.User) *model.User {
	return &model.User{
		ID:        data.Uid,
		Username:  data.Username,
		Password:  data.Password,
		CreatedAt: strconv.FormatInt(data.CreatedAt.Unix(), 10),
		UpdatedAt: strconv.FormatInt(data.UpdatedAt.Unix(), 10),
	}
}

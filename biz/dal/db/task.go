package db

import (
	"context"
	"time"
	"work3/pkg/constants"
)

func CreateTask(ctx context.Context, title, content string, uid, startAt, endAt int64) (*Task, error) {

	var taskResp *Task

	taskResp = &Task{
		Uid:     uid,
		Title:   title,
		Content: content,
		StartAt: time.Unix(startAt, 0),
		EndAt:   time.Unix(endAt, 0),
	}

	err := DB.
		WithContext(ctx).
		Table(constants.TaskTable).
		Create(&taskResp).
		Error

	if err != nil {
		return nil, err
	}

	return taskResp, nil
}

func UpdateTask(ctx context.Context, id, uid, status int64) (err error) {

	if id == 0 {
		err = DB.
			WithContext(ctx).
			Table(constants.TaskTable).
			Where("uid = ?", uid).
			Update("status", status).
			Error
	} else {
		err = DB.
			WithContext(ctx).
			Table(constants.TaskTable).
			Where("uid = ?", uid).
			Where("id = ?", id).
			Update("status", status).
			Error
	}

	return

}

func DeleteOneTask(ctx context.Context, uid, id int64) (err error) {
	err = DB.
		WithContext(ctx).
		Table(constants.TaskTable).
		Delete(&Task{
			Uid: uid,
			Id:  id,
		}).
		Error
	return
}

func DeleteTask(ctx context.Context, uid, status int64) (err error) {

	if status == 2 {
		err = DB.
			WithContext(ctx).
			Table(constants.TaskTable).
			Delete(&Task{Uid: uid}).
			Error

	} else {
		err = DB.
			WithContext(ctx).
			Table(constants.TaskTable).
			Delete(&Task{
				Uid:    uid,
				Status: status,
			}).
			Error
	}

	return
}

func QueryTaskListByStatus(ctx context.Context, pagesize, pagenum, uid, status int64) ([]*Task, int64, error) {

	var taskResp []*Task
	var err error
	var count int64

	if status == 2 {
		err = DB.
			WithContext(ctx).
			Table(constants.TaskTable).
			Where("uid=?", uid).
			Limit(int(pagesize)).
			Offset(int((pagenum - 1) * pagesize)).
			Count(&count).
			Find(&taskResp).
			Error

	} else {
		err = DB.
			WithContext(ctx).
			Table(constants.TaskTable).
			Where("uid=?", uid).
			Where("status=?", status).
			Limit(int(pagesize)).
			Offset(int((pagenum - 1) * pagesize)).
			Count(&count).
			Find(&taskResp).
			Error
	}

	if err != nil {
		return nil, -1, err
	}

	return taskResp, count, nil
}

func QueryTaskListByKeyWord(ctx context.Context, pageSize, pageNum, uid int64, keyWord string) ([]*Task, int64, error) {

	var taskResp []*Task
	var count int64

	err := DB.WithContext(ctx).
		Table(constants.TaskTable).
		Where("uid = ?", uid).
		Where("title LIKE ?", keyWord).
		Limit(int(pageSize)).
		Offset(int((pageNum - 1) * pageSize)).
		Order("created_at desc").
		Count(&count).
		Find(&taskResp).
		Error

	if err != nil {
		return nil, -1, err
	}

	return taskResp, count, nil
}

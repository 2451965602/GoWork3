package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"work3/biz/dal/db"
	"work3/biz/model/task"
)

type TaskService struct {
	ctx context.Context
	c   *app.RequestContext
}

func NewTaskService(ctx context.Context, c *app.RequestContext) *TaskService {
	return &TaskService{ctx: ctx, c: c}
}

func (s *TaskService) AddTask(req *task.AddTaskRequest) (*db.Task, error) {
	return db.CreateTask(s.ctx, req.Title, req.Content, GetUidFormContext(s.c), req.StartAt, req.EndAt)
}

func (s *TaskService) UpdateTask(req *task.UpdateTaskRequest) error {
	return db.UpdateTask(s.ctx, req.ID, GetUidFormContext(s.c), req.Status)
}

func (s *TaskService) DeleteOneTask(req *task.DeleteTaskRequest) error {
	return db.DeleteOneTask(s.ctx, GetUidFormContext(s.c), req.ID)
}

func (s *TaskService) DeleteTaskByStatus(req *task.DeleteTaskByStatusRequest) error {
	return db.DeleteTask(s.ctx, GetUidFormContext(s.c), req.Status)
}

func (s *TaskService) QueryTaskListByStatus(req *task.QueryTaskListByStatusRequest) ([]*db.Task, int64, error) {
	return db.QueryTaskListByStatus(s.ctx, req.PageSize, req.PageNum, GetUidFormContext(s.c), req.Status)
}

func (s *TaskService) QueryTaskListByKeyWord(req *task.QueryTaskListByKeywordRequest) ([]*db.Task, int64, error) {
	return db.QueryTaskListByKeyWord(s.ctx, req.PageSize, req.PageNum, GetUidFormContext(s.c), req.Keyword)
}

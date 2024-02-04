// Code generated by hertz generator.

package task

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"work3/biz/pack"
	"work3/biz/service"

	"github.com/cloudwego/hertz/pkg/app"
	"work3/biz/model/task"
)

// AddTask .
// @router /task/add [POST]
func AddTask(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.AddTaskRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(task.AddTaskResponse)

	taskResp, err := service.NewTaskService(ctx, c).AddTask(&req)

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Data = pack.Task(taskResp)

	pack.SendResponse(c, resp, consts.StatusOK)
}

// UpdateTask .
// @router /task/update [PUT]
func UpdateTask(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.UpdateTaskRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(task.UpdateTaskResponse)

	err = service.NewTaskService(ctx, c).UpdateTask(&req)

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	pack.SendResponse(c, resp, consts.StatusOK)
}

// DeleteTask .
// @router /task/delete [DELETE]
func DeleteTask(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.DeleteTaskRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(task.DeleteTaskResponse)

	err = service.NewTaskService(ctx, c).DeleteOneTask(&req)

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)

	pack.SendResponse(c, resp, consts.StatusOK)
}

// DeleteTaskByStatus .
// @router /task/delete/status [DELETE]
func DeleteTaskByStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.DeleteTaskByStatusRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(task.DeleteTaskByStatusResponse)

	err = service.NewTaskService(ctx, c).DeleteTaskByStatus(&req)

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)

	pack.SendResponse(c, resp, consts.StatusOK)
}

// QueryTaskByStatus .
// @router /task/querylist/status [GET]
func QueryTaskByStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.QueryTaskListByStatusRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(task.QueryTaskListBystatusResponse)

	list, count, err := service.NewTaskService(ctx, c).QueryTaskListByStatus(&req)

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Data = pack.TaskList(list, count)

	pack.SendResponse(c, resp, consts.StatusOK)
}

// QueryTaskListByKeyword .
// @router /task/querylist/keyword [GET]
func QueryTaskListByKeyword(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.QueryTaskListByKeywordRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp := new(task.QueryTaskListByKeywordResponse)

	list, count, err := service.NewTaskService(ctx, c).QueryTaskListByKeyWord(&req)

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Data = pack.TaskList(list, count)

	pack.SendResponse(c, resp, consts.StatusOK)
}

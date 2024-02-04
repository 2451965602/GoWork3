namespace go task

include "model.thrift"

//增
struct AddTaskRequest{
    1: required string title, //标题
    2: required string content //内容
    3: required i64 start_at, //开始时间
    4: required i64 end_at, //结束时间
}

struct AddTaskResponse{
    1: model.BaseResp base,
    2: model.Task data
}

//改
struct UpdateTaskRequest{
    1: required i64 id, //日程id   0-所有，
    2: required i64 status, //状态 0-未完成，1-完成
}

struct UpdateTaskResponse{
    1: model.BaseResp base,
}

//删
struct DeleteTaskRequest{
    1: required i64 id,//日程id
}

struct DeleteTaskResponse{
    1: model.BaseResp base,
}


struct DeleteTaskByStatusRequest{
    1: required i64 status,//状态 0-未完成，1-完成，2-所有
}

struct DeleteTaskByStatusResponse {
    1: model.BaseResp base,
}
//查

struct QueryTaskListByStatusRequest{
    1:required i64 page_size,//单页个数
    2:required i64 page_num,//页数
    3:required i64 status, //状态 0-未完成，1-完成，2-所有
}

struct QueryTaskListBystatusResponse{
    1:model.BaseResp base;
    2:model.TaskList data
}

struct QueryTaskListByKeywordRequest{
    1:required i64 page_size, //单页个数
    2:required i64 page_num, //页数
    3:required string keyword, //关键字
}

struct QueryTaskListByKeywordResponse{
    1:model.BaseResp base,
    2:model.TaskList data,
}

service TaskService{
    AddTaskResponse AddTask(1:AddTaskRequest req) (api.post="/task/add"),

    UpdateTaskResponse UpdateTask(1:UpdateTaskRequest req) (api.put="/task/update"),

    DeleteTaskResponse DeleteTask(1:DeleteTaskRequest req) (api.delete="/task/delete"),
    DeleteTaskByStatusResponse DeleteTaskByStatus(1:DeleteTaskByStatusRequest req) (api.delete="/task/delete/status"),

    QueryTaskListBystatusResponse QueryTaskByStatus(1:QueryTaskListByStatusRequest req) (api.get="/task/querylist/status"),
    QueryTaskListByKeywordResponse QueryTaskListByKeyword(1:QueryTaskListByKeywordRequest req) (api.get="/task/querylist/keyword"),
}
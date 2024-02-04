namespace go model

struct User {
    1:required i64 id,
    2:required string username,
    3:required string password,
    4:required string created_at,
    5:required string updated_at,
}

struct Task {
    1:required i64 id,
    2:required string title,
    3:required string content,
    4:required i64 status,
    5:required string created_at,
    6:required string updated_at
    7:required string start_at
    8:required string end_at
}

struct TaskList{
    1:required list<Task> items,
    2:required i64 total,
}

struct BaseResp{
    1:required i64 code,
    2:required string msg,
}
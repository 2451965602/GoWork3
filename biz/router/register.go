// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	model "work3/biz/router/model"
	task "work3/biz/router/task"
	user "work3/biz/router/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	user.Register(r)

	model.Register(r)

	task.Register(r)

}

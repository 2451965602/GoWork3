// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"work3/biz/middleware"
	"work3/biz/pack"
	"work3/biz/service"

	"github.com/cloudwego/hertz/pkg/app"
	"work3/biz/model/user"
)

// Register .
// @router /user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}
	resp := new(user.RegisterResponse)

	userResp, err := service.NewUserService(ctx, c).Register(&req)

	if err != nil {
		pack.SendFailResponse(c, err)
		return
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Data = pack.User(userResp)
	pack.SendResponse(c, resp, consts.StatusOK)
}

// Login .
// @router /user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	middleware.JwtMiddleware.LoginHandler(ctx, c)
}

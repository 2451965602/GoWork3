package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"log"
	"net/http"
	"time"
	"work3/biz/dal/db"
	"work3/biz/model/user"
	"work3/biz/pack"
)

var (
	IdentityKey   = "uid"
	JwtMiddleware *jwt.HertzJWTMiddleware
)

func InitJwt() {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:         "ToDOList",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		TokenHeadName: "Bearer",
		IdentityKey:   IdentityKey,

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*db.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.Uid,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[IdentityKey]
		},

		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			pack.SendResponse(c, message, code)
		},

		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {

			c.JSON(http.StatusOK, map[string]interface{}{
				"code":   http.StatusOK,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
				"msg":    "登录成功",
			})
		},

		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct user.LoginRequest
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			users, err := db.LoginCheck(ctx, loginStruct.Username, loginStruct.Password)
			if err != nil {
				return nil, err
			}
			return users, nil
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}

func Init() {
	InitJwt()
	errInit := JwtMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}

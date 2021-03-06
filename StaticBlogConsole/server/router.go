package server

import (
	"os"
	"sbc/api"
	"sbc/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	//中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	//r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		authed := v1.Group("/")
		authed.Use(middleware.AuthRequired())
		{
			// User Routing
			authed.POST("user/create", api.UserCreateMD)
			authed.DELETE("user/logout", api.UserLogout)
		}
	}

	return r
}

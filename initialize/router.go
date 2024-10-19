package initialize

import (
	"github.com/gin-gonic/gin"
	"pandora-server/global"
	"pandora-server/middleware"
	"pandora-server/route"
)

// 路由初始化
func Router() *gin.Engine {
	// 初始化一个没有任何中间件的路由引擎
	r := gin.New()

	// 中间件
	r.Use(middleware.AccessLog)       // 请求日志中间件
	r.Use(middleware.Cors)            // 跨域访问中间件
	r.Use(middleware.Exception)       // 异常捕获中间件
	auth, err := middleware.JWTAuth() // JWT 认证中间件
	if err != nil {
		panic("JWT 认证中间件初始化失败：" + err.Error())
	}

	// 路由组
	route.PublicRoutes(r.Group(global.SystemApiPrefix), auth)                        // 免登录路由组
	route.PublicAuthRoutes(r.Group(global.SystemApiPrefix), auth)                    // 登录路由组
	route.SystemUserAuthRoutes(r.Group(global.SystemApiPrefix+"/system/user"), auth) // 系统用户路由组
	route.SystemMenuAuthRoutes(r.Group(global.SystemApiPrefix+"/system/menu"), auth) // 系统菜单路由组

	return r
}

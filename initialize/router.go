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
	r.Use(middleware.Cors)      // 跨域访问中间件
	r.Use(middleware.Exception) // 异常捕获中间件

	// 路由组
	rg := r.Group(global.SystemApiPrefix)
	route.PublicRoutes(rg)
	return r
}

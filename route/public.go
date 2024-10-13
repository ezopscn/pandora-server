package route

import (
	"github.com/gin-gonic/gin"
	"pandora-server/api/v1"
)

// 开放路由组
func PublicRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/health", v1.HealthHandler) // 健康检查接口
	rg.GET("/info", v1.InfoHandler)     // 系统信息接口
	return rg
}

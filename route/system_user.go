package route

import (
	"github.com/gin-gonic/gin"
	"pandora-server/api/v1"
)

// 用户路由
func SystemUserRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/list", v1.GetSystemUserListHandler)                        // 获取用户列表接口
	rg.GET("/detail", v1.GetSystemUserDetailHandler)                    // 获取当前用户详情接口
	rg.GET("/specified/detail", v1.GetSystemUserSpecifiedDetailHandler) // 获取指定用户详情接口
	return rg
}

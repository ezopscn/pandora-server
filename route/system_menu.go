package route

import (
	"github.com/gin-gonic/gin"
	"pandora-server/api/v1"
)

// 菜单路由
func SystemMenuRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/tree", v1.GetSystemMenuTreeHandler) // 系统菜单树接口
	return rg
}

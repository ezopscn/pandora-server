package route

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"pandora-server/api/v1"
)

// 菜单路由
func SystemMenuAuthRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	authRG := rg.Use(auth.MiddlewareFunc())
	authRG.GET("/tree", v1.GetSystemMenuTreeHandler) // 系统菜单树接口
	return authRG
}

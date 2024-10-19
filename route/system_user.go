package route

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"pandora-server/api/v1"
)

// 用户路由
func SystemUserAuthRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	authRG := rg.Use(auth.MiddlewareFunc())
	authRG.GET("/list", v1.GetSystemUserListHandler)                                              // 获取用户列表接口
	authRG.GET("/detail", v1.GetSystemUserDetailHandler)                                          // 获取当前用户详情接口
	authRG.GET("/specified/detail/by/id/:id", v1.GetSystemUserSpecifiedDetailHandler)             // 通过用户 Id 获取指定用户详情接口
	authRG.GET("/specified/detail/by/username/:username", v1.GetSystemUserSpecifiedDetailHandler) // 通过用户名获取指定用户详情接口
	authRG.GET("/specified/detail/by/phone/:phone", v1.GetSystemUserSpecifiedDetailHandler)       // 通过手机号获取指定用户详情接口
	authRG.GET("/specified/detail/by/email/:email", v1.GetSystemUserSpecifiedDetailHandler)       // 通过邮箱获取指定用户详情接口
	return authRG
}

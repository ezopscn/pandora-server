package route

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"pandora-server/api/v1"
)

// 开放路由组
func PublicRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/health", v1.HealthHandler)  // 健康检查接口
	rg.GET("/info", v1.InfoHandler)      // 系统信息接口
	rg.POST("/login", auth.LoginHandler) // 用户登录
	return rg
}

// 登录路由组
func PublicAuthRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	authRG := rg.Use(auth.MiddlewareFunc())
	authRG.GET("/logout", auth.LogoutHandler) // 用户注销登录
	return authRG
}

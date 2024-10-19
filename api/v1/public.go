package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora-server/global"
	"pandora-server/pkg/response"
)

// 健康检测接口
func HealthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

// 系统信息接口
func InfoHandler(ctx *gin.Context) {
	response.SuccessWithData(gin.H{
		"enName":         global.SystemENName,
		"cnName":         global.SystemCNName,
		"describe":       global.SystemDescribe,
		"version":        global.SystemVersion,
		"go":             global.SystemGoVersion,
		"developerName":  global.SystemDeveloperName,
		"developerEmail": global.SystemDeveloperEmail,
	})
}

// 用户登录接口
func LoginHandler(ctx *gin.Context) {
	response.Success()
}

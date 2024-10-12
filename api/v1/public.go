package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pandora-server/global"
)

// 健康检测接口
func HealthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

// 系统信息接口
func InfoHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"enName":         global.SystemENName,
		"cnName":         global.SystemCNName,
		"describe":       global.SystemDescribe,
		"version":        global.SystemVersion,
		"go":             global.SystemGoVersion,
		"developerName":  global.SystemDeveloperName,
		"developerEmail": global.SystemDeveloperEmail,
	})
}

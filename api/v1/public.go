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
		"enName":         global.SYSTEM_EN_NAME,
		"cnName":         global.SYSTEM_CN_NAME,
		"describe":       global.SYSTEM_DESCRIBE,
		"version":        global.SystemVersion,
		"go":             global.SYSTEM_GO_VERSION,
		"developerName":  global.SYSTEM_DEVELOPER_NAME,
		"developerEmail": global.SYSTEM_DEVELOPER_EMAIL,
	})
}

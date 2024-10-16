package v1

import (
	"github.com/gin-gonic/gin"
	"pandora-server/pkg/response"
	"pandora-server/service"
)

// 获取用户列表
func GetSystemUserListHandler(ctx *gin.Context) {
	systemUsers, err := service.SearchSystemUserListService(ctx)
	if err != nil {
		response.FailedWithMessage("查询用户列表失败")
		return
	}
	response.SuccessWithData(systemUsers)
}

// 获取当前用户的详细信息
func GetSystemUserDetailHandler(ctx *gin.Context) {
	response.Success()
}

// 获取指定用户的详细信息
func GetSystemUserSpecifiedDetailHandler(ctx *gin.Context) {
	response.Success()
}

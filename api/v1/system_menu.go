package v1

import (
	"github.com/gin-gonic/gin"
	"pandora-server/pkg/response"
	"pandora-server/service"
)

// 菜单树列表接口
func GetSystemMenuTreeHandler(ctx *gin.Context) {
	tree, err := service.GenerateSystemMenuTreeByRoleIdService(0)
	if err != nil {
		response.FailedWithMessage("获取菜单树数据失败")
		return
	}
	response.SuccessWithData(tree)
}

package v1

import (
	"github.com/gin-gonic/gin"
	"pandora-server/pkg/response"
	"pandora-server/pkg/trans"
	"pandora-server/pkg/utils"
	"pandora-server/service"
	"strconv"
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
	// 从 ctx 中解析出当前用户的相关信息
	username, err := utils.ExtractStringResultFromContext(ctx, "username")
	if err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 查询用户信息
	systemUser, err := service.SearchSystemUserDetailService("username", username)
	if err != nil {
		response.FailedWithMessage("查询当前用户的详细信息失败")
		return
	}
	response.SuccessWithData(systemUser)
}

// 获取指定用户的详细信息
func GetSystemUserSpecifiedDetailHandler(ctx *gin.Context) {
	// 解析 URI 参数
	var key string
	var value interface{}

	// 判断用户传递的参数是啥
	idStr := ctx.Param("id")
	username := ctx.Param("username")
	phone := ctx.Param("phone")
	email := ctx.Param("email")

	if idStr != "" {
		v, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil || v <= 0 {
			response.FailedWithMessage("获取查询用户 id 失败")
			return
		}
		key = "id"
		value = v
	} else if username != "" {
		key = "username"
		value = username
	} else if phone != "" {
		if !utils.IsPhoneNumber(phone) {
			response.FailedWithMessage("获取查询用户的手机号失败")
			return
		}
		key = "phone"
		value = phone
	} else if email != "" {
		if !utils.IsEmail(email) {
			response.FailedWithMessage("获取查询用户的邮箱失败")
			return
		}
		key = "email"
		value = email
	} else {
		response.FailedWithMessage("获取查询参数失败")
		return
	}

	// 查询用户
	systemUser, err := service.SearchSystemUserDetailService(key, value)
	if err != nil {
		response.FailedWithMessage("查询指定用户的详细信息失败")
		return
	}

	// 通过判断是否隐藏联系方式调整数据
	if systemUser.HidePhone == trans.Uint(1) {
		systemUser.Phone = utils.HidePhoneNumber(systemUser.Phone)
	}
	response.SuccessWithData(systemUser)
}

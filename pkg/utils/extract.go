package utils

import (
	"errors"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 从 Context 中提取 Uint 类型指定数据
func ExtractUintResultFromContext(ctx *gin.Context, keyword string) (value uint, err error) {
	claims := jwt.ExtractClaims(ctx)
	v, ok := claims[keyword].(float64) // 注意客户端请求过来的 JSON 会变成 float64 类型
	if !ok {
		err = errors.New(fmt.Sprintf("从 Context 提取 %s 的值失败", keyword))
	}
	value = uint(v)
	return
}

// 从 Context 中提取 String 类型指定数据
func ExtractStringResultFromContext(ctx *gin.Context, keyword string) (value string, err error) {
	claims := jwt.ExtractClaims(ctx)
	value, ok := claims[keyword].(string)
	if !ok {
		err = errors.New(fmt.Sprintf("从 Context 提取 %s 的值失败", keyword))
	}
	return
}

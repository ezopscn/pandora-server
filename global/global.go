package global

import (
	"embed"
)

var (
	// 打包的静态资源文件用于全局使用
	FS embed.FS
)

const (
	// 系统名称
	SystemName = "pandora"
	// 系统英文名称
	SystemENName = "Pandora"
	// 系统中文名称
	SystemCNName = "潘多拉"
	// 系统描述
	SystemDescribe = "一个由 Go + React 开发的 Kubernetes 运维管理系统"
	// Go 版本
	SystemGoVersion = "1.23.0"
	// 开发者
	SystemDeveloperName = "DK"
	// 开发者邮箱
	SystemDeveloperEmail = "ezops.cn@gmail.com"
	// 版本文件
	SystemVersionFile = "version"
	// API 前缀
	SystemApiPrefix = "/api/v1"
)

var (
	// 系统版本
	SystemVersion = ""
	// 监听地址
	SystemListenAddress = ""
	// 监听端口
	SystemListenPort = ""
	// 配置文件
	SystemConfigFilename = "pandora.yaml"
)

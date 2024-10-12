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
	SYSTEM_NAME = "pandora"
	// 系统英文名称
	SYSTEM_EN_NAME = "Pandora"
	// 系统中文名称
	SYSTEM_CN_NAME = "潘多拉"
	// 系统描述
	SYSTEM_DESCRIBE = "一个由 Go + React 开发的 Kubernetes 运维管理系统"
	// Go 版本
	SYSTEM_GO_VERSION = "1.23.0"
	// 开发者
	SYSTEM_DEVELOPER_NAME = "DK"
	// 开发者邮箱
	SYSTEM_DEVELOPER_EMAIL = "ezops.cn@gmail.com"
	// 版本文件
	SYSTEM_VERSION_FILE = "version"
	// API 前缀
	SYSTEM_API_PREFIX = "/api/v1"
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

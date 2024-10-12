package main

import (
	"embed"
	"pandora-server/cmd"
	"pandora-server/global"
)

//go:embed pandora.yaml
//go:embed version
var fs embed.FS // Go 1.16 版本之后提供的将静态资源打包的方法，写法固定，可以将目录也打包

func main() {
	// 配置静态资源全局使用
	global.FS = fs

	// 读取版本信息文件
	version, err := global.FS.ReadFile(global.SystemVersionFile)
	if err != nil || string(version) == "" {
		panic("版本信息文件读取失败，请确认项目中 version 文件是否真实存在或者配置正确：" + err.Error())
	} else {
		global.SystemVersion = string(version)
	}

	// 程序入口
	cmd.Execute()
}

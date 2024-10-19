package initialize

import (
	"fmt"
	"pandora-server/global"
	"pandora-server/initialize/data"
	"time"
)

// 系统基础数据初始化
func MigrateData() {
	fmt.Println("开始进行系统基础数据初始化：" + time.Now().Format(global.MillisecondTimeFormat))
	data.InitializeSystemUser() // 用户
	data.InitializeSystemRole() // 角色
	data.InitializeSystemMenu() // 菜单
	fmt.Println("系统基础数据初始化完成：" + time.Now().Format(global.MillisecondTimeFormat))
}

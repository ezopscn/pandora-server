package initialize

import (
	"fmt"
	"pandora-server/global"
	"time"
)

// 系统基础数据初始化
func MigrateData() {
	fmt.Println("开始进行系统基础数据初始化：" + time.Now().Format(global.MillisecondTimeFormat))
	fmt.Println("系统基础数据初始化完成：" + time.Now().Format(global.MillisecondTimeFormat))
}

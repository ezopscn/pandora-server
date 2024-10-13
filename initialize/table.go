package initialize

import (
	"fmt"
	"pandora-server/global"
	"pandora-server/model"
	"time"
)

// 数据结构同步
func MigrateTable() {
	fmt.Println("开始进行数据结构（数据表）同步：" + time.Now().Format(global.MillisecondTimeFormat))
	_ = global.MySQLDB.AutoMigrate(new(model.SystemUser)) // 用户
	fmt.Println("数据结构（数据表）同步完成：" + time.Now().Format(global.MillisecondTimeFormat))
}

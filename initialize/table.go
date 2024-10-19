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
	_ = global.MySQLDB.AutoMigrate(new(model.SystemUser))    // 用户
	_ = global.MySQLDB.AutoMigrate(new(model.SystemRole))    // 角色
	_ = global.MySQLDB.AutoMigrate(new(model.SystemMenu))    // 菜单
	_ = global.MySQLDB.AutoMigrate(new(model.SystemApiType)) // 接口类型
	_ = global.MySQLDB.AutoMigrate(new(model.SystemApi))     // 接口详情
	fmt.Println("数据结构（数据表）同步完成：" + time.Now().Format(global.MillisecondTimeFormat))
}

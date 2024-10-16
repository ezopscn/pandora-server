package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"pandora-server/global"
	"time"
)

// MySQL 连接初始化
func MySQL() {
	// 数据库连接串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&collation=%s&timeout=%dms&%s",
		global.Config.MySQL.Username,
		global.Config.MySQL.Password,
		global.Config.MySQL.Host,
		global.Config.MySQL.Port,
		global.Config.MySQL.Database,
		global.Config.MySQL.Charset,
		global.Config.MySQL.Collation,
		global.Config.MySQL.Timeout,
		global.Config.MySQL.ExtraParam)

	// 连接数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn, // 数据库连接字符串
		DefaultStringSize: 170, // varchar 默认长度，太长影响查询
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
			// TablePrefix:   "tb_", // 表名前缀
		},
		DisableForeignKeyConstraintWhenMigrating: true,  // 禁用外键
		IgnoreRelationshipsWhenMigrating:         false, // 开启会导致 many2many 的表创建失败
		QueryFields:                              true,  // 解决查询索引失效问题
	})

	// 错误处理
	if err != nil {
		panic("MySQL 连接初始化失败：" + err.Error())
	}

	// 设置数据库连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(global.Config.MySQL.MaxOpenConns)
	sqlDB.SetMaxIdleConns(global.Config.MySQL.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(time.Duration(global.Config.MySQL.MaxIdleTime) * time.Minute)

	// 设置全局数据库连接，方便后续使用
	global.MySQLDB = db
}

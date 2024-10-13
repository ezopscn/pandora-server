package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pandora-server/global"
	"pandora-server/initialize"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(tableCmd)
	migrateCmd.AddCommand(dataCmd)
}

// 迁移命令
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "数据初始化，可以通过 --help 查看具体用法",
}

// 迁移表
var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "数据库结构（数据表）初始化",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(global.LOGO)
		initialize.Config()
		initialize.MySQL()
		initialize.MigrateTable()
	},
}

// 迁移数据
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "系统基础依赖数据初始化",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(global.LOGO)
		initialize.Config()
		initialize.MySQL()
		initialize.MigrateData()
	},
}

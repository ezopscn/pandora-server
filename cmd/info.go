package cmd

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"pandora-server/global"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

// 系统信息命令
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "查看当前程序的相关信息",
	Run: func(cmd *cobra.Command, args []string) {
		t := table.NewWriter()
		t.AppendHeader(table.Row{"名称", "说明"})
		t.AppendRows([]table.Row{
			{"英文名称", global.SYSTEM_EN_NAME},
			{"中文名称", global.SYSTEM_CN_NAME},
			{"项目说明", global.SYSTEM_DESCRIBE},
			{"系统版本", global.SystemVersion},
			{"GO 版本", global.SYSTEM_GO_VERSION},
			{"开发人员", global.SYSTEM_DEVELOPER_NAME},
			{"开发邮箱", global.SYSTEM_DEVELOPER_EMAIL},
		})
		fmt.Println(t.Render())
	},
}

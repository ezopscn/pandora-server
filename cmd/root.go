package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pandora-server/global"
)

// 命令入口
var rootCmd = &cobra.Command{
	Use:   global.SystemName,
	Short: fmt.Sprintf("%s（%s），%s", global.SystemENName, global.SystemCNName, global.SystemDescribe),
	// 如果有相关的 action 要执行，请取消下面这行代码的注释
	// Run: func(cmd *cobra.Command, args []string) { },
}

// 所有子命令添加到 root 命令，输入 cmd 的入口
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic("程序异常：" + err.Error())
	}
}

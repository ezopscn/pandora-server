package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
	"pandora-server/global"
	"pandora-server/initialize"
	"pandora-server/pkg/utils"
	"time"
)

func init() {
	rootCmd.AddCommand(startCmd)
	// 指定配置文件参数
	startCmd.Flags().StringVarP(&global.SystemConfigFilename, "config", "", global.SystemConfigFilename, "可选，指定服务启动配置文件")
	startCmd.Flags().StringVarP(&global.SystemListenAddress, "listen", "", global.SystemListenAddress, "可选，指定服务启动的监听地址")
	startCmd.Flags().StringVarP(&global.SystemListenPort, "port", "", global.SystemListenPort, "可选，指定服务启动的监听端口")
}

// 启动命令
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "参数化启动服务，可以通过 --help 查看更多参数用法",
	Run: func(cmd *cobra.Command, args []string) {
		// Logo
		fmt.Println(global.LOGO)

		// 配置文件初始化
		initialize.Config()

		// 判断用户是否传递监听地址是否合法，没有传参，则使用配置文件中的配置
		if global.SystemListenAddress != "" {
			if !utils.IsIPv4(global.SystemListenAddress) {
				panic("命令行参数传递的监听地址不合法")
			}
		} else {
			if !utils.IsIPv4(global.Config.System.Listen) {
				panic("配置文件设置的监听地址不合法")
			}
			global.SystemListenAddress = global.Config.System.Listen
		}

		// 判断用户是否传递监听端口是否合法，没有传参，则使用配置文件中的配置
		if global.SystemListenPort != "" {
			if !utils.IsPort(global.SystemListenPort) {
				panic("命令行参数传递的监听端口不合法")
			}
		} else {
			if !utils.IsPort(global.Config.System.Port) {
				panic("配置文件设置的监听端口不合法")
			}
			global.SystemListenPort = global.Config.System.Port
		}

		// 初始化日志
		initialize.SystemLogger() // 系统日志初始化
		initialize.AccessLogger() // 访问日志初始化

		// 路由初始化
		r := initialize.Router()
		addr := fmt.Sprintf("%s:%s", global.SystemListenAddress, global.SystemListenPort)
		server := http.Server{
			Addr:    addr,
			Handler: r,
		}

		// 启动服务
		go func() {
			err := server.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				panic("服务启动异常：" + err.Error())
			}
		}()

		// 服务启动信息
		fmt.Println("服务启动完成，监听地址：" + addr)

		// 接收优雅关闭信号
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		// 等待5秒然后停止服务
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			panic("服务停止异常：" + err.Error())
		}
		fmt.Println("服务停止成功")
	},
}

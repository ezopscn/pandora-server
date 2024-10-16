package initialize

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"pandora-server/global"
	"pandora-server/pkg/utils"
)

// 配置文件初始化
func Config() {
	// 读取的数据
	var bs []byte
	var err error

	// Viper 读取文件
	v := viper.New()
	v.SetConfigType("yaml")

	// 优先读取本地文件，然后才是 embed 打包的配置
	filename := global.SystemConfigFilename
	exist := utils.FileExist(filename)
	if exist {
		fmt.Println("开始加载指定的配置文件:", filename)
		bs, err = os.ReadFile(filename)
	} else {
		fmt.Println("开始加载内置的配置文件:", filename)
		bs, err = global.FS.ReadFile(filename)
	}
	if err != nil {
		panic("配置文件加载失败：" + err.Error())
	}

	// 解析配置
	err = v.ReadConfig(bytes.NewReader(bs))
	if err != nil {
		panic("配置文件解析失败：" + err.Error())
	}

	// 将配置设置到内存中，方便后续调用
	settings := v.AllSettings()
	for i, setting := range settings {
		v.Set(i, setting)
	}

	// 设置全局变量，方便调用
	err = v.Unmarshal(&global.Config)
	if err != nil {
		panic("配置信息全局设置失败：" + err.Error())
	}
}

package main

import (
	"fmt"
	"goweb/config"
	"goweb/router"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	// 初始化配置
	config.InitConfig()
	fmt.Printf("===== 服务：%s启动成功 =====\n", config.AppConfig.App.Name)

	// 设置路由分组
	ginServer := router.SetupRouter()
	// 设置服务端口号
	_ = ginServer.Run(":" + config.AppConfig.App.Port)
}

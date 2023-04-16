/**
 * @Author: woshishabii
 * @Description:
 * @File: main
 * @Version: 0.0.1
 * @Date: 4/15/2023 3:19 PM
 */

package main

import (
	"GoZhixue/conf"
	"GoZhixue/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载陆游
	r := server.NewRouter()
	r.Run(":3000")
}

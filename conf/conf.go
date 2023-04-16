/**
 * @Author: woshishabii
 * @Description:
 * @File: conf
 * @Version: 0.0.1
 * @Date: 4/15/2023 4:22 PM
 */

package conf

import (
	"GoZhixue/cache"
	"GoZhixue/model"
	"GoZhixue/util"
	"github.com/joho/godotenv"
	"os"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 链接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}

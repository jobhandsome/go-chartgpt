package gin

import (
	"github.com/gin-gonic/gin"
	"go-chatgpt/app/frontend/v1/config"
)

// Init /** 初始化Gin
func Init() *gin.Engine {
	// 设置环境模式：DebugMode debug模式  ReleaseMode 生产模式  TestMode 测试模式
	switch config.Config.Debug {
	case true:
		gin.SetMode(gin.DebugMode)
	case false:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	// 创建一个默认路由
	r := gin.Default()
	// 注册路由
	InitRoutes(r)

	return r
}

package gin

import (
	"github.com/gin-gonic/gin"
	"go-chatgpt/app/frontend/v1/config"
	"go-chatgpt/app/middleware"
	"go-chatgpt/pkg/logger"
	"go-chatgpt/pkg/orm"
	"go.uber.org/zap"
	"log"
)

// Init /** 初始化Gin
func Init() *gin.Engine {
	// 设置环境模式：DebugMode debug模式  ReleaseMode 生产模式  TestMode 测试模式
	switch config.Config.Debug {
	case true:
		gin.SetMode(gin.DebugMode)
		// 自定义路由日志的格式
		gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
			log.Printf("[GIN-debug] - [%v] - [%v] - [%v] (%v handles)\n", httpMethod, absolutePath, handlerName, nuHandlers)
		}
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	// 引入 gorm'
	orm.Init()

	// 初始化 logger
	logger.Init()

	// 创建一个默认路由
	r := gin.Default()

	// 初始化中间件
	r.Use(middleware.GinLogger(zap.L()), middleware.GinRecovery(zap.L(), true))

	// 注册路由
	InitRoutes(r)

	return r
}

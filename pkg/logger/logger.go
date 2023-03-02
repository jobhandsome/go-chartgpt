package logger

import (
	"go-chatgpt/pkg/utils"
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Init initlogger 初始化日志
func Init() {
	// 1、encoder
	encndercfg := zap.NewProductionEncoderConfig()
	encndercfg.TimeKey = "time"                          // 改变时间的key
	encndercfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 更改时间格式
	encndercfg.EncodeLevel = zapcore.CapitalLevelEncoder //将日志级别大写并带有颜色
	enconder := zapcore.NewJSONEncoder(encndercfg)

	// 2、writerSyncer 将日志写到文件和终端
	file, _ := os.OpenFile(path.Dir(path.Dir(utils.GetCurrentAbPathByCaller()))+"/logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fileWS := zapcore.AddSync(file)
	consoleWS := zapcore.AddSync(os.Stdout)

	// 3、设置loglevel
	level := zapcore.DebugLevel

	// 创建zapcore
	core := zapcore.NewCore(enconder, zapcore.NewMultiWriteSyncer(fileWS, consoleWS), level)
	// 创建logger
	logger := zap.New(core)

	// 替换zap全局的logger
	zap.ReplaceGlobals(logger)
	zap.L().Info(" logger init success")

}

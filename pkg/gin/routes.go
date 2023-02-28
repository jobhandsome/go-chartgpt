package gin

import (
	"github.com/gin-gonic/gin"
	frontendV1 "go-chatgpt/app/frontend/v1/route"
)

type Option func(engine *gin.Engine)

var Options []Option

// Register /** 注册路由配置
func Register(opts ...Option) {
	Options = append(Options, opts...)
}

func InitRoutes(r *gin.Engine) {
	// 注册路由
	Register(frontendV1.Router)
	// 加载注册路由
	for _, opt := range Options {
		opt(r)
	}
}

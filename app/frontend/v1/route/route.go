package route

import (
	"github.com/gin-gonic/gin"
	"go-chatgpt/app/frontend/v1/controller"
)

func Router(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/test", controller.Test)
		v1.GET("/completions", controller.GetCompletions)
		// 用户组
		user := v1.Group("/user")
		{
			user.POST("/login", controller.UserLogin)
			user.POST("/register", controller.UserRegister)
			user.GET("/wechat-code", controller.WechatCode)
		}
	}
}

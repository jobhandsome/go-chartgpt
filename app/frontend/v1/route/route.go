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
	}
}

package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-chatgpt/app/frontend/v1/config"
	"net/http"
)

func Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    config.Config,
	})
	return
}

func GetCompletions(ctx *gin.Context) {
	// 提问内容
	text := "写一个golang的快速排序"
	resp, err := Completions(text)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	respMap := make(map[string]any)
	_ = json.Unmarshal(resp, &respMap)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    respMap,
	})
	return
}

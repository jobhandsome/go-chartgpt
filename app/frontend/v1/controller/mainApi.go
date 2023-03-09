package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-chatgpt/app/frontend/types"
	"go-chatgpt/pkg/utils"
	"net/http"
)

func Test(ctx *gin.Context) {
	ApiKey := GetApiKey()

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    ApiKey,
	})
	return
}

func GetCompletions(ctx *gin.Context) {
	// 提问内容
	var params types.TestRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		utils.ServerError(ctx, err.Error())
		return
	}

	fmt.Println(params.Text)

	if len(params.Text) == 0 {
		utils.ServerError(ctx, "参数错误")
		return
	}
	resp, err := Completions(params.Text)
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

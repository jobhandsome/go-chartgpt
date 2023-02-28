package main

import (
	"go-chatgpt/app/frontend/v1/config"
	"go-chatgpt/pkg/gin"
	"go-chatgpt/pkg/yaml"
	"strconv"
)

func main() {
	// 初始化 yaml
	yaml.Init()
	r := gin.Init()
	_ = r.Run(":" + strconv.Itoa(config.Config.Port))
}

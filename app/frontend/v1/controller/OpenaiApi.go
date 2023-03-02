package controller

import (
	"bytes"
	"encoding/json"
	"go-chatgpt/pkg/utils"
	"go.uber.org/zap"
)

const (
	ApiUrl      = "https://api.openai.com/v1/" // OpenAI 地址
	ApiKey      = "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"
	ContentType = "application/json"
)

// Completions /** 根据提示生成关联段落或文字
// text 指定提示文字
func Completions(text string) ([]byte, error) {
	url := ApiUrl + "completions"
	// 设置header头个
	headers := map[string]string{
		"Content-Type":  ContentType,
		"Authorization": "Bearer " + ApiKey,
	}

	// 组装body数据
	formData := map[string]any{
		"model":       "text-davinci-003",
		"prompt":      text,
		"temperature": 0,
		"max_tokens":  4000,
	}
	jsonData, _ := json.Marshal(formData)
	resp, err := utils.SendRequest(url, bytes.NewBuffer(jsonData), headers, "POST")
	zap.L().Info(string(resp))
	if err != nil {
		return nil, err
	}
	return resp, err
}

// ImageGenerations 根据提示生成图像
// text 指定提示文字
// n 生成的图像数
// size 图像尺寸 256x256, 512x512, or 1024x1024
// responseFormat 返回格式 url / b64_json
func ImageGenerations(text string, n int, size string, responseFormat string) ([]byte, error) {
	url := ApiUrl + "images/generations"
	// 设置header头个
	headers := map[string]string{
		"Content-Type":  ContentType,
		"Authorization": "Bearer " + ApiKey,
	}

	// 组装body数据
	formData := map[string]any{
		"prompt":          text,
		"n":               n,
		"size":            size,
		"response_format": responseFormat,
	}

	jsonData, _ := json.Marshal(formData)
	resp, err := utils.SendRequest(url, bytes.NewBuffer(jsonData), headers, "POST")
	zap.L().Info(string(resp))
	if err != nil {
		return nil, err
	}
	return resp, err
}

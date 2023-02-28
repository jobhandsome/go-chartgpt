package utils

import (
	"errors"
	"io"
	"net/http"
	"time"
)

// SendRequest /** 发送request
func SendRequest(url string, body io.Reader, addHeaders map[string]string, method string) (resp []byte, err error) {
	// 1、创建req
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	// 2、设置headers
	if len(addHeaders) > 0 {
		for k, v := range addHeaders {
			req.Header.Add(k, v)
		}
	} else {
		req.Header.Add("Content-Type", "application/json")
	}

	// 3、发送http请求
	client := &http.Client{Timeout: 60 * time.Second}
	response, err := client.Do(req)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	if response.StatusCode != 200 {
		err = errors.New("http status err")
		return
	}

	// 4、结果读取
	resp, err = io.ReadAll(response.Body)
	return
}

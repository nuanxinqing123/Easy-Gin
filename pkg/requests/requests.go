package requests

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	"Easy-Gin/config"
)

// Request HTTP请求客户端结构体
type Request struct {
	client *resty.Client // resty客户端实例
}

// New 创建一个新的HTTP请求客户端
func New() *Request {
	client := resty.New()

	// 设置DeBug模式
	// 当应用程序运行在debug模式时，开启HTTP请求的调试输出
	if config.GinConfig.App.Mode == "debug" {
		client.SetDebug(true)
	}

	// 设置请求重试机制
	client.SetRetryCount(3)                  // 最大重试次数为3次
	client.SetRetryWaitTime(1 * time.Second) // 重试等待时间为1秒

	return &Request{
		client: client,
	}
}

// Get 发送GET请求
// url: 请求地址
// params: URL查询参数
func (r *Request) Get(url string, params map[string]string) (*resty.Response, error) {
	// 记录请求日志
	config.GinLOG.Debug("GET: " + url)
	config.GinLOG.Debug(fmt.Sprintf("params: %v", params))

	// 发送GET请求
	return r.client.R().
		SetQueryParams(params).
		Get(url)
}

// Post 发送POST请求
// url: 请求地址
// body: 请求体数据
func (r *Request) Post(url string, body any) (*resty.Response, error) {
	// 记录请求日志
	config.GinLOG.Debug("POST: " + url)
	config.GinLOG.Debug(fmt.Sprintf("body: %v", body))

	// 发送POST请求
	return r.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)
}

// Delete 发送DELETE请求
// url: 请求地址
// body: 请求体数据
func (r *Request) Delete(url string, body any) (*resty.Response, error) {
	// 记录请求日志
	config.GinLOG.Debug("DELETE: " + url)
	config.GinLOG.Debug(fmt.Sprintf("body: %v", body))

	// 发送DELETE请求
	return r.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Delete(url)
}

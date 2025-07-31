package main

import (
	"mingdemo/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		// 使用自定义的请求核心处理函数
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}

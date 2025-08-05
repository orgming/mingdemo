package main

import (
	"mingdemo/framework"
	"mingdemo/framework/middleware"
	"net/http"
)

func main() {
	core := framework.NewCore()

	core.Use(middleware.Test1(), middleware.Test2())

	registerRouter(core)
	server := &http.Server{
		// 使用自定义的请求核心处理函数
		Handler: core,
		Addr:    ":8888",
	}
	server.ListenAndServe()
}

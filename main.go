// Copyright 2025 Andy Ron. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	mingHttp "github.com/orgming/mingdemo/app/http"
	"github.com/orgming/mingdemo/app/provider/demo"
	"github.com/orgming/mingdemo/framework/gin"
	"github.com/orgming/mingdemo/framework/middleware"
	"github.com/orgming/mingdemo/framework/provider/app"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建engine结构
	core := gin.New()
	// 绑定具体的服务
	core.Bind(&app.MingAppProvider{})
	core.Bind(&demo.DemoProvider{})

	core.Use(gin.Recovery())
	core.Use(middleware.Cost())

	mingHttp.Routes(core)

	server := &http.Server{
		// 使用自定义的请求核心处理函数
		Handler: core,
		Addr:    ":8888",
	}

	// 这个 Goroutine 是启动服务的 Goroutine
	go func() {
		server.ListenAndServe()
	}()

	// 当前的 Goroutine 等待信号量
	quit := make(chan os.Signal)
	//  监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前 Goroutine 等待信号
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}

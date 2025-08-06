package main

import (
	"context"
	"log"
	"mingdemo/framework"
	"mingdemo/framework/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	core := framework.NewCore()

	//core.Use(middleware.Test1(), middleware.Test2())
	core.Use(middleware.Recovery())
	core.Use(middleware.Cost())

	registerRouter(core)
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

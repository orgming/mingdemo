package http

import "github.com/orgming/ming/framework/gin"

// NewHttpEngine 创建了一个绑定了路由的Web引擎
func NewHttpEngine() (*gin.Engine, error) {
	// 设置为Release，为的是默认在启动中不输出调试信息
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	Routes(r)

	return r, nil
}

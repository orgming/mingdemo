package framework

import "net/http"

// 框架核心结构
type Core struct {
}

// 框架核心结构初始化
func NewCore() *Core {
	return &Core{}
}

// 实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
}

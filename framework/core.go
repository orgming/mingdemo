package framework

import (
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	//router map[string]ControllerHandler
	router map[string]map[string]ControllerHandler // 二级map
}

// NewCore 框架核心结构初始化
func NewCore() *Core {
	// 定义二级map
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}

	// 把二级map写入一级map
	router := map[string]map[string]ControllerHandler{}
	router["GET"] = getRouter
	router["POST"] = postRouter
	router["PUT"] = putRouter
	router["DELETE"] = deleteRouter

	return &Core{router: router}
}

// === http method

// 路由注册，将路由注册函数按照 Method 名拆分为 4 个方法：Get、Post、Put 和 Delete。
// 统一将路由转为大写。

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router["GET"][strings.ToUpper(url)] = handler
}

func (c *Core) Post(url string, handler ControllerHandler) {
	c.router["POST"][strings.ToUpper(url)] = handler
}

func (c *Core) Put(url string, handler ControllerHandler) {
	c.router["PUT"][strings.ToUpper(url)] = handler
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	c.router["DELETE"][strings.ToUpper(url)] = handler
}

// === http method end

// ServeHTTP 实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 封装自定义context
	ctx := NewContext(request, response)

	// 寻找路由
	router := c.FindRouteByRequest(request)
	if router == nil {
		// 没有找到路由就打印日志
		ctx.Json(404, "Not Found")
		return
	}

	// 调用路由函数，如果err说明内部错误，返回500
	if err := router(ctx); err != nil {
		ctx.Json(500, "Inner Error")
		return
	}
}

// FindRouteByRequest  匹配路由方法
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	uri := strings.ToUpper(request.URL.Path)
	method := strings.ToUpper(request.Method)
	if methodHandlers, ok := c.router[method]; ok {
		if handler, ok := methodHandlers[uri]; ok {
			return handler
		}
	}
	return nil
}

package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	//router map[string]ControllerHandler
	//router map[string]map[string]ControllerHandler // 二级map
	router map[string]*Tree // 改为trie树
}

// NewCore 框架核心结构初始化
func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()

	return &Core{router: router}
}

// === http method

// 路由注册，将路由注册函数按照 Method 名拆分为 4 个方法：Get、Post、Put 和 Delete。
// 统一将路由转为大写。

func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
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
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

// Group 在core中初始化Group
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

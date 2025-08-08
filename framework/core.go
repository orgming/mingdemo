// Copyright 2025 Andy Ron. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core 框架核心结构
type Core struct {
	router      map[string]*Tree // 改为trie树
	middlewares []ControllerHandler
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

// Use 注册中间件
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

// === http method

// 路由注册，将路由注册函数按照Method名拆分为4个方法：Get、Post、Put 和 Delete。
// 统一将路由转为大写。

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	// 将core的middleware 和 handlers结合起来
	all := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, all); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	all := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, all); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	all := append(c.middlewares, handlers...)
	if err := c.router["PUT"].AddRouter(url, all); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	all := append(c.middlewares, handlers...)
	if err := c.router["DELETE"].AddRouter(url, all); err != nil {
		log.Fatal("add router error: ", err)
	}
}

// === http method end

// ServeHTTP 实现Handler接口。所有请求都进入这个函数, 这个函数负责路由分发
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 封装自定义context
	ctx := NewContext(request, response)

	// 寻找路由
	node := c.FindRouteNodeByRequest(request)
	if node == nil {
		// 没有找到路由就打印日志
		ctx.SetStatus(404).Json("Not Found")
		return
	}

	ctx.SetHandlers(node.handlers)

	// 设置路由参数
	params := node.parseParamsFromEndNode(request.URL.Path)
	ctx.SetParams(params)

	// 调用路由函数，如果err说明内部错误，返回500
	if err := ctx.Next(); err != nil {
		ctx.SetStatus(500).Json("Inner Error")
		return
	}
}

// FindRouteNodeByRequest  匹配路由方法
func (c *Core) FindRouteNodeByRequest(request *http.Request) *node {
	uri := strings.ToUpper(request.URL.Path)
	method := strings.ToUpper(request.Method)
	if methodHandlers, ok := c.router[method]; ok {
		return methodHandlers.root.matchNode(uri)
	}
	return nil
}

// Group 在core中初始化Group
func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

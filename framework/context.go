package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

// Context 代表当前请求上下文
type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	ctx            context.Context
	handlers       []ControllerHandler // 当前请求的handler链条
	index          int                 // 当前请求调用到调用链的哪个节点

	// 是否超时标记位
	hasTimeout bool
	// 写保护机制
	writeMux *sync.Mutex

	// 路由参数
	params map[string]string
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writeMux:       &sync.Mutex{},
		index:          -1,
	}
}

// #region base function

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writeMux
}

func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// #endregion

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

// #region implement context.Context

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key any) any {
	return ctx.BaseContext().Value(key)
}

// #endregion

func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

// Next 调用context的下一个函数 【核心函数】
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

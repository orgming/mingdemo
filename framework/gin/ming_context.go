// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"context"
	"github.com/orgming/mingdemo/framework"
)

func (c *Context) BaseContext() context.Context {
	return c.Request.Context()
}

// engine 实现 container 的绑定封装

func (engine *Engine) Bind(provider framework.ServiceProvider) error {
	return engine.container.Bind(provider)
}

func (engine *Engine) IsBind(key string) bool {
	return engine.container.IsBind(key)
}

// context 实现 container 的几个封装

func (c *Context) Make(key string) (any, error) {
	return c.container.Make(key)
}

func (c *Context) MustMake(key string) any {
	return c.container.MustMake(key)
}

func (c *Context) MakeNew(key string, params []any) (any, error) {
	return c.container.MakeNew(key, params)
}

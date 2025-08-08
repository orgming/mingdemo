// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"github.com/spf13/cast"
)

// IRequest 代表请求包含的方法
type IRequest interface {

	// 从请求地址 url 中获取参数，如 foo.com?a=1&b=bar&c[]=bar

	MyQueryInt(key string, def int) (int, bool)
	MyQueryInt64(key string, def int64) (int64, bool)
	MyQueryFloat32(key string, def float32) (float32, bool)
	MyQueryFloat64(key string, def float64) (float64, bool)
	MyQueryBool(key string, def bool) (bool, bool)
	MyQueryString(key string, def string) (string, bool)
	MyQueryStringSlice(key string, def []string) ([]string, bool)
	MyQuery(key string) any

	// 从路由匹配中获取参数，如 /book/:id
	MyParamInt(key string, def int) (int, bool)
	MyParamInt64(key string, def int64) (int64, bool)
	MyParamFloat32(key string, def float32) (float32, bool)
	MyParamFloat64(key string, def float64) (float64, bool)
	MyParamBool(key string, def bool) (bool, bool)
	MyParamString(key string, def string) (string, bool)
	MyParamStringSlice(key string, def []string) ([]string, bool)
	MyParam(key string) any

	// form表单中带的参数
	MyFormInt(key string, def int) (int, bool)
	MyFormInt64(key string, def int64) (int64, bool)
	MyFormFloat32(key string, def float32) (float32, bool)
	MyFormFloat64(key string, def float64) (float64, bool)
	MyFormBool(key string, def bool) (bool, bool)
	MyFormString(key string, def string) (string, bool)
	MyFormStringSlice(key string, def []string) ([]string, bool)
	MyForm(key string) any

	BindJson(obj any) error
	BindXml(obj any) error
	// 其他格式
	GetRawData() ([]byte, error)

	// 基础信息
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	Headers() map[string][]string
	Header(key string) (string, bool)

	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

// #region url中参数

// QueryAll 获取请求地址中所有参数
func (c *Context) QueryAll() map[string][]string {
	c.initQueryCache()
	return map[string][]string(c.queryCache)
}

func (c *Context) MyQuery(key string) any {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		return vals[0]
	}
	return nil
}

// MyQueryInt 获取Int类型的请求参数
func (c *Context) MyQueryInt(key string, def int) (int, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) MyQueryInt64(key string, def int64) (int64, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) MyQueryFloat64(key string, def float64) (float64, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) MyQueryFloat32(key string, def float32) (float32, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) MyQueryBool(key string, def bool) (bool, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) MyQueryString(key string, def string) (string, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return def, false
}

func (c *Context) MyQueryStringSlice(key string, def []string) ([]string, bool) {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

// #endregion

// #region 路由参数

func (c *Context) MyParam(key string) any {
	if val, ok := c.Params.Get(key); ok {
		return val
	}
	return nil
}

func (c *Context) MyParamInt(key string, def int) (int, bool) {
	if val := c.MyParam(key); val != nil {
		return cast.ToInt(val), true
	}
	return def, false
}

func (c *Context) MyParamInt64(key string, def int64) (int64, bool) {
	if val := c.MyParam(key); val != nil {
		return cast.ToInt64(val), true
	}
	return def, false
}

func (c *Context) MyParamFloat64(key string, def float64) (float64, bool) {
	if val := c.MyParam(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return def, false
}

func (c *Context) MyParamFloat32(key string, def float32) (float32, bool) {
	if val := c.MyParam(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return def, false
}

func (c *Context) MyParamBool(key string, def bool) (bool, bool) {
	if val := c.MyParam(key); val != nil {
		return cast.ToBool(val), true
	}
	return def, false
}

func (c *Context) MyParamString(key string, def string) (string, bool) {
	if val := c.MyParam(key); val != nil {
		return cast.ToString(val), true
	}
	return def, false
}

// #endregion

// #region 表单中参数

func (c *Context) FormAll() map[string][]string {
	c.initFormCache()
	return map[string][]string(c.formCache)
}

func (c *Context) MyFormInt(key string, def int) (int, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) MyFormInt64(key string, def int64) (int64, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) MyFormFloat64(key string, def float64) (float64, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) MyFormFloat32(key string, def float32) (float32, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) MyFormBool(key string, def bool) (bool, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return def, false
}

func (c *Context) MyFormString(key string, def string) (string, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		return vals[0], true
	}
	return def, false
}

func (c *Context) MyFormStringSlice(key string, def []string) ([]string, bool) {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return def, false
}

func (c *Context) MyForm(key string) any {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0]
		}
	}
	return nil
}

// #endregion

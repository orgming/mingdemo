// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

// IResponse 代表返回包含的方法
type IResponse interface {
	// Json 输出
	IJson(obj any) IResponse

	// Jsonp 输出
	IJsonp(obj any) IResponse

	// xml 输出
	IXml(obj any) IResponse

	// html 输出
	IHtml(template string, obj any) IResponse

	// string
	IText(format string, values ...any) IResponse

	// 重定向
	IRedirect(path string) IResponse

	// header
	ISetHeader(key string, val string) IResponse

	// Cookie
	ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	ISetStatus(code int) IResponse

	// 设置 200 状态
	ISetOkStatus() IResponse
}

func (c *Context) IJsonp(obj any) IResponse {
	callbackFunc := c.Query("callback")
	c.ISetHeader("Content-Type", "application/javascript")
	// 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成 XSS 攻击
	callback := template.JSEscapeString(callbackFunc) // TODO

	// 输出函数名
	_, err := c.Writer.Write([]byte(callback))
	if err != nil {
		return c
	}
	// 输出左括号
	_, err = c.Writer.Write([]byte("("))
	if err != nil {
		return c
	}
	// 数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return c
	}
	_, err = c.Writer.Write(ret)
	if err != nil {
		return c
	}
	// 输出右括号
	_, err = c.Writer.Write([]byte(")"))
	if err != nil {
		return c
	}
	return c
}

func (c *Context) IJson(obj any) IResponse {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return c.ISetStatus(http.StatusInternalServerError)
	}
	c.ISetHeader("Content-Type", "application/json")
	c.Writer.Write(bytes)
	return nil
}

func (c *Context) IXml(obj any) IResponse {
	bytes, err := xml.Marshal(obj)
	if err != nil {
		return c.ISetStatus(http.StatusInternalServerError)
	}
	c.ISetHeader("Content-Type", "application/html")
	c.Writer.Write(bytes)
	return nil
}

func (c *Context) IHtml(file string, obj any) IResponse {
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return c
	}
	// 将obj和模版进行结合
	if err := t.Execute(c.Writer, obj); err != nil { // TODO
		return c
	}

	c.ISetHeader("Content-Type", "application/html")
	return c
}

func (c *Context) IText(format string, values ...any) IResponse {
	out := fmt.Sprintf(format, values...)
	c.ISetHeader("Content-type", "application/text")
	c.Writer.Write([]byte(out))
	return c
}

func (c *Context) IRedirect(path string) IResponse {
	http.Redirect(c.Writer, c.Request, path, http.StatusMovedPermanently)
	return c
}

func (c *Context) ISetHeader(key string, val string) IResponse {
	c.Writer.Header().Set(key, val)
	return c
}

func (c *Context) ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
		SameSite: 1, // TODO
	})
	return c
}

func (c *Context) ISetStatus(code int) IResponse {
	c.Writer.WriteHeader(code)
	return c
}

func (c *Context) ISetOkStatus() IResponse {
	c.ISetStatus(http.StatusOK)
	return c
}

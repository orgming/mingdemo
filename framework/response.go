package framework

import (
	"encoding/json"
	"html/template"
	"net/http"
)

// IResponse 代表返回包含的方法
type IResponse interface {
	// Json 输出
	Json(obj any) IResponse

	// Jsonp 输出
	Jsonp(obj any) IResponse

	// xml 输出
	Xml(obj any) IResponse

	// html 输出
	Html(template string, obj any) IResponse

	// string
	Text(format string, values ...any) IResponse

	// 重定向
	Redirect(path string) IResponse

	// header
	SetHeader(key string, val string) IResponse

	// Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	SetStatus(code int) IResponse

	// 设置 200 状态
	SetOkStatus() IResponse
}

func (ctx *Context) Jsonp(obj any) IResponse {
	callbackFunc, _ := ctx.QueryString("callbak", "callback_function")
	ctx.SetHeader("Content-Type", "application/javascript")
	// 输出到前端页面的时候需要注意下进行字符过滤，否则有可能造成 XSS 攻击
	callbak := template.JSEscapeString(callbackFunc) // TODO

	// 输出函数名
	_, err := ctx.responseWriter.Write([]byte(callbak))
	if err != nil {
		return ctx
	}
	// 输出左括号
	_, err = ctx.responseWriter.Write([]byte("("))
	if err != nil {
		return ctx
	}
	// 数据函数参数
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}
	_, err = ctx.responseWriter.Write(ret)
	if err != nil {
		return ctx
	}
	// 输出右括号
	_, err = ctx.responseWriter.Write([]byte(")"))
	if err != nil {
		return ctx
	}
	return ctx
}

func (ctx *Context) Json(obj any) IResponse {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return ctx.SetStatus(http.StatusInternalServerError)
	}
	ctx.SetHeader("Content-Type", "application/json")
	ctx.responseWriter.Write(bytes)
	return nil
}

func (Context) Xml(obj any) IResponse {
	//TODO implement me
	panic("implement me")
}

func (ctx *Context) Html(file string, obj any) IResponse {
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	// 将obj和模版进行结合
	if err := t.Execute(ctx.responseWriter, obj); err != nil { // TODO
		return ctx
	}

	ctx.SetHeader("Content-Type", "application/html")
	return ctx
}

func (ctx *Context) Text(format string, values ...any) IResponse {
	//TODO implement me
	panic("implement me")
}

func (ctx *Context) Redirect(path string) IResponse {
	//TODO implement me
	panic("implement me")
}

func (ctx *Context) SetHeader(key string, val string) IResponse {
	//TODO implement me
	panic("implement me")
}

func (ctx *Context) SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse {
	//TODO implement me
	panic("implement me")
}

func (ctx *Context) SetStatus(code int) IResponse {
	//TODO implement me
	panic("implement me")
}

func (ctx *Context) SetOkStatus() IResponse {
	//TODO implement me
	panic("implement me")
}

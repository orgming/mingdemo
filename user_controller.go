package main

import (
	"github.com/orgming/mingdemo/framework/gin"
	"time"
)

func UserLoginController(c *gin.Context) {
	foo, _ := c.MyQueryString("foo", "def")
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)
	c.ISetOkStatus().IJson("Ok, UserLoginController: " + foo)
}

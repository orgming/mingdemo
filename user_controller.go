package main

import (
	"mingdemo/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	// 等待10s才结束执行
	time.Sleep(10 * time.Second)
	c.SetOkStatus().Json("Ok, UserLoginController: " + foo)
	return nil
}

package main

import "mingdemo/framework"

func UserLoginController(c *framework.Context) error {
	// 打印控制器名字
	c.SetOkStatus().Json("Ok, UserLoginController")
	return nil
}

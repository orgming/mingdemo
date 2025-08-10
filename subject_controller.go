package main

import (
	"github.com/orgming/mingdemo/framework/gin"
	"github.com/orgming/mingdemo/provider/demo"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("Ok, SubjectAddController")
}

// 对应路由 /subject/list/all
func SubjectListController(c *gin.Context) {

	// 获取Demo服务实例
	demoService := c.MustMake(demo.Key).(demo.Service)
	// 调用服务实例的方法
	foo := demoService.GetFoo()

	c.ISetOkStatus().IJson(foo)
}

func SubjectDeleteController(c *gin.Context) {
	c.ISetOkStatus().IJson("Ok, SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("Ok, SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	c.ISetOkStatus().IJson("Ok, SubjectGetController")
}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("Ok, SubjectNameController")
}

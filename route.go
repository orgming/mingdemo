package main

import (
	"github.com/orgming/mingdemo/framework"
	"github.com/orgming/mingdemo/framework/middleware"
)

// registerRouter 注册路由规则
func registerRouter(core *framework.Core) {
	// 需求1，2：HTTP方法，静态路由匹配
	// 在core中使用middleware.Test3() 为单个路由增加中间件
	core.Get("/user/login", middleware.Test3(), UserLoginController)

	// 需求3：批量通用前缀
	subjectApi := core.Group("/subject")
	subjectApi.Use(middleware.Test3())
	{
		// 需求4：动态路由
		// 在group中使用middleware.Test3() 为单个路由增加中间件
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Delete("/:id", SubjectDeleteController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}

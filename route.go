package main

import "mingdemo/framework"

// registerRouter 注册路由规则
func registerRouter(core *framework.Core) {
	//core.Get("foo", FooControllerHandler)
	// 需求1，2：HTTP方法，静态路由匹配
	core.Get("/user/login", UserLoginController)

	// 需求3：批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 需求4：动态路由
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Delete("/:id", SubjectDeleteController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}

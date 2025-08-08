package main

import (
	"github.com/orgming/mingdemo/framework/gin"
)

// registerRouter 注册路由规则
func registerRouter(core *gin.Engine) {
	// 需求1，2：HTTP方法，静态路由匹配
	core.GET("/user/login", UserLoginController)

	// 需求3：批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 需求4：动态路由
		// 在group中使用middleware.Test3() 为单个路由增加中间件
		subjectApi.GET("/:id", SubjectGetController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.DELETE("/:id", SubjectDeleteController)
		subjectApi.GET("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", SubjectNameController)
		}

	}
}

package http

import (
	"github.com/orgming/ming/app/http/module/demo"
	"github.com/orgming/ming/framework/gin"
	"github.com/orgming/ming/framework/middleware/static"
)

func Routes(r *gin.Engine) {
	//r.Static("/dist/", "./dist/")

	// /路径先去./dist目录下查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))

	demo.Register(r)
}

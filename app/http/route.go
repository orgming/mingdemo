package http

import (
	"github.com/orgming/mingdemo/app/http/module/demo"
	"github.com/orgming/mingdemo/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")
	demo.Register(r)
}

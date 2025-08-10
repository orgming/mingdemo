package http

import (
	"github.com/orgming/ming/app/http/module/demo"
	"github.com/orgming/ming/framework/gin"
)

func Routes(r *gin.Engine) {
	r.Static("/dist/", "./dist/")
	demo.Register(r)
}

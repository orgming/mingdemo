package main

import "mingdemo/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}

// Copyright 2025 Andy Ron. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/orgming/ming/app/console"
	"github.com/orgming/ming/app/http"
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/provider/app"
	"github.com/orgming/ming/framework/provider/kernel"
)

func main() {
	// 初始化服务容器
	container := framework.NewMingContainer()

	// 绑定APP服务提供者
	container.Bind(&app.MingAppProvider{})

	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	if engine, err := http.NewHttpEngine(); err == nil {
		container.Bind(&kernel.MingKernelProvider{
			HttpEngine: engine,
		})
	}

	//
	console.RunCommand(container)

}

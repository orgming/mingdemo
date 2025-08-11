package contract

import "net/http"

// 把创建 Engine 的过程封装为一个服务接口协议

// KernelKey 提供kernel服务凭证
const KernelKey = "ming:k"

// Kernel 接口提供框架最核心的结构
type Kernel interface {
	// HttpEngine http.Handler结构，作为net/http框架使用, 实际上是gin.Engine
	HttpEngine() http.Handler
}

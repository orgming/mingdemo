package contract

// AppKey 定义字符串凭证
const AppKey = "ming:app"

// App 定义接口
type App interface {
	// Version 当前版本
	Version() string
	// BaseFolder 项目基础目录
	BaseFolder() string
	// ConfigFolder 配置文件目录
	ConfigFolder() string
	// LogFolder 日志目录
	LogFolder() string
	// ProviderFolder 业务自己的服服务提供目录
	ProviderFolder() string
	// MiddlewareFolder 业务自己的中间件目录
	MiddlewareFolder() string
	// CommandFolder 业务定义的命令
	CommandFolder() string
	// RuntimeFolder 业务的运行中间态信息
	RuntimeFolder() string
	// TestFolder 测试所需要的信息
	TestFolder() string
}

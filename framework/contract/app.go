package contract

// AppKey 定义字符串凭证
const AppKey = "ming:app"

// App 应用目录服务接口
type App interface {
	// Version 当前版本
	Version() string
	// BaseFolder 应用根目录
	BaseFolder() string
	// ConfigFolder 配置文件目录
	ConfigFolder() string
	// LogFolder 日志目录
	LogFolder() string
	// ProviderFolder 业务的服务提供者和对应接口的存放目录
	ProviderFolder() string
	// MiddlewareFolder 业务自己的中间件目录
	MiddlewareFolder() string
	// CommandFolder 业务定义的命令
	CommandFolder() string
	// RuntimeFolder 业务的运行中间态信息（进程ID等）
	RuntimeFolder() string
	// TestFolder 测试所需要的信息
	TestFolder() string
}

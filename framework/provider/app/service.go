package app

import (
	"errors"
	"github.com/google/uuid"
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/util"
	"path/filepath"
)

// MingApp 代表ming框架的App实现
type MingApp struct {
	container  framework.Container // 服务容器
	baseFolder string              // 基础路径
	appID      string

	configMap map[string]string // 配置加载
}

func (app MingApp) AppID() string {
	return app.appID
}

func (app MingApp) Version() string {
	return "0.0.1"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (app MingApp) BaseFolder() string {
	if app.baseFolder != "" {
		return app.baseFolder
	}
	// 没有设置就是用命令行参数
	//var baseFolder string
	//flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数，默认为当前路径")
	//flag.Parse()
	//if baseFolder != "" {
	//	return baseFolder
	//}

	return util.GetExecDirectory()
}

func (app MingApp) ConfigFolder() string {
	if val, ok := app.configMap["config_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "config")
}

func (app MingApp) LogFolder() string {
	if val, ok := app.configMap["log_folder"]; ok {
		return val
	}
	return filepath.Join(app.StorageFolder(), "log")
}

func (app MingApp) HttpFolder() string {
	if val, ok := app.configMap["http_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "http")
}

func (app MingApp) ConsoleFolder() string {
	if val, ok := app.configMap["console_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "console")
}

func (app MingApp) StorageFolder() string {
	if val, ok := app.configMap["storage_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "storage")
}

func (app MingApp) ProviderFolder() string {
	if val, ok := app.configMap["provider_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "provider")
}

func (app MingApp) MiddlewareFolder() string {
	if val, ok := app.configMap["middleware_folder"]; ok {
		return val
	}
	return filepath.Join(app.HttpFolder(), "middleware")
}

func (app MingApp) CommandFolder() string {
	if val, ok := app.configMap["command_folder"]; ok {
		return val
	}
	return filepath.Join(app.ConsoleFolder(), "command")
}

func (app MingApp) RuntimeFolder() string {
	if val, ok := app.configMap["runtime_folder"]; ok {
		return val
	}
	return filepath.Join(app.StorageFolder(), "runtime")
}

func (app MingApp) TestFolder() string {
	if val, ok := app.configMap["test_folder"]; ok {
		return val
	}
	return filepath.Join(app.BaseFolder(), "test")
}

func NewMingApp(params ...any) (any, error) {
	if len(params) != 2 {
		return nil, errors.New("params must be 2")
	}
	// 两个参数一个是容器，一个是基础路径
	container := params[0].(framework.Container)
	baseFolder := params[1].(string)
	appID := uuid.New().String()
	return &MingApp{
		container:  container,
		baseFolder: baseFolder,
		appID:      appID,
	}, nil
}

func (app *MingApp) LoadAppConfig(kv map[string]string) {
	for key, val := range kv {
		app.configMap[key] = val
	}
}

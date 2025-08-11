package app

import (
	"errors"
	"flag"
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
}

func (m MingApp) Version() string {
	return "0.0.1"
}

// BaseFolder 表示基础目录，可以代表开发场景的目录，也可以代表运行时候的目录
func (m MingApp) BaseFolder() string {
	if m.baseFolder != "" {
		return m.baseFolder
	}
	// 没有设置就是用命令行参数
	var baseFolder string
	flag.StringVar(&baseFolder, "base_folder", "", "base_folder参数，默认为当前路径")
	flag.Parse()
	if baseFolder != "" {
		return baseFolder
	}

	return util.GetExecDirectory()
}

func (m MingApp) ConfigFolder() string {
	return filepath.Join(m.BaseFolder(), "config")
}

func (m MingApp) LogFolder() string {
	return filepath.Join(m.BaseFolder(), "log")
}
func (m MingApp) ProviderFolder() string {
	return filepath.Join(m.BaseFolder(), "provider")
}

func (m MingApp) MiddlewareFolder() string {
	return filepath.Join(m.BaseFolder(), "middleware")
}

func (m MingApp) CommandFolder() string {
	return filepath.Join(m.BaseFolder(), "command")
}

func (m MingApp) RuntimeFolder() string {
	return filepath.Join(m.BaseFolder(), "runtime")
}

func (m MingApp) TestFolder() string {
	return filepath.Join(m.BaseFolder(), "test")
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

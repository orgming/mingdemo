package app

import (
	"github.com/orgming/mingdemo/framework"
	"github.com/orgming/mingdemo/framework/contract"
)

// MingAppProvider 提供App的具体实现方法
type MingAppProvider struct {
	BaseFolder string
}

func (m *MingAppProvider) Register(container framework.Container) framework.NewInstance {
	return NewMingApp
}

// Boot 启动调用
func (m *MingAppProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (m *MingAppProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (m *MingAppProvider) Params(container framework.Container) []any {
	return []any{container, m.BaseFolder}
}

func (m *MingAppProvider) Name() string {
	return contract.AppKey
}

package config

import (
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
	"path/filepath"
)

type MingConfigProvider struct{}

func (p *MingConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewMingConfig
}

func (p *MingConfigProvider) Boot(c framework.Container) error {
	return nil
}

func (p *MingConfigProvider) IsDefer() bool {
	return false
}

func (p *MingConfigProvider) Params(c framework.Container) []any {
	appService := c.MustMake(contract.AppKey).(contract.App)
	envService := c.MustMake(contract.EnvKey).(contract.Env)
	return []any{
		c,
		filepath.Join(appService.ConfigFolder(), envService.AppEnv()),
		envService.All(),
	}
}

func (p *MingConfigProvider) Name() string {
	return contract.ConfigKey
}

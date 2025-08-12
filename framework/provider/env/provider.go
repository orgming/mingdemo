package env

import (
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
)

type MingEnvProvider struct {
	Folder string
}

func (p *MingEnvProvider) Register(c framework.Container) framework.NewInstance {
	return NewMingEnv
}

func (p *MingEnvProvider) Boot(c framework.Container) error {
	app := c.MustMake(contract.AppKey).(contract.App)
	p.Folder = app.BaseFolder()
	return nil
}

func (p *MingEnvProvider) IsDefer() bool {
	return false
}

func (p *MingEnvProvider) Params(c framework.Container) []any {
	return []any{p.Folder}
}

func (p *MingEnvProvider) Name() string {
	return contract.EnvKey
}

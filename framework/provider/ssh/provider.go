package ssh

import (
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
)

type SSHProvider struct {
}

func (p *SSHProvider) Register(c framework.Container) framework.NewInstance {
	return NewMingSSH
}

func (p *SSHProvider) Boot(c framework.Container) error {
	return nil
}

func (p *SSHProvider) IsDefer() bool {
	return true
}

func (p *SSHProvider) Params(c framework.Container) []any {
	return []any{c}
}

func (p *SSHProvider) Name() string {
	return contract.SSHKey
}

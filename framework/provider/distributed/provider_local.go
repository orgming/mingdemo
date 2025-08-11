package distributed

import (
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
)

type LocalDistributedProvider struct {
}

func (p *LocalDistributedProvider) Register(c framework.Container) framework.NewInstance {
	return NewLocalDistributedService
}

func (p *LocalDistributedProvider) Boot(c framework.Container) error {
	return nil
}

func (p *LocalDistributedProvider) IsDefer() bool {
	return false
}

func (p *LocalDistributedProvider) Params(container framework.Container) []any {
	return []any{container}
}

func (p *LocalDistributedProvider) Name() string {
	return contract.DistributedKey
}

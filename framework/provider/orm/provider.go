package orm

import (
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
)

type GormProvider struct {
}

func (p GormProvider) Register(container framework.Container) framework.NewInstance {
	return NewMingGorm
}

func (p GormProvider) Boot(container framework.Container) error {
	return nil
}

func (p GormProvider) IsDefer() bool {
	return true
}

func (p GormProvider) Params(container framework.Container) []any {
	return []any{container}
}

func (p GormProvider) Name() string {
	return contract.ORMKey
}

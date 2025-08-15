package trace

import (
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
)

type MingTraceProvider struct {
	c framework.Container
}

func (p *MingTraceProvider) Register(c framework.Container) framework.NewInstance {
	return NewMingTraceService
}

func (p *MingTraceProvider) Boot(c framework.Container) error {
	p.c = c
	return nil
}

func (p *MingTraceProvider) IsDefer() bool {
	return false
}

func (p *MingTraceProvider) Params(c framework.Container) []any {
	return []any{p.c}
}

func (p *MingTraceProvider) Name() string {
	return contract.TraceKey
}

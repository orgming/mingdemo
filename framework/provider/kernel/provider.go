package kernel

import (
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
	"github.com/orgming/ming/framework/gin"
)

// MingKernelProvider 提供web引擎
type MingKernelProvider struct {
	HttpEngine *gin.Engine
}

func (p *MingKernelProvider) Register(c framework.Container) framework.NewInstance {
	return NewMingKernelService
}

func (p *MingKernelProvider) Boot(c framework.Container) error {
	if p.HttpEngine == nil {
		p.HttpEngine = gin.Default()
	}
	p.HttpEngine.SetContainer(c)
	return nil
}

func (p *MingKernelProvider) IsDefer() bool {
	return false
}

func (p *MingKernelProvider) Params(c framework.Container) []any {
	return []any{p.HttpEngine}
}

func (p *MingKernelProvider) Name() string {
	return contract.KernelKey
}

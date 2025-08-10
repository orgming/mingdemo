package demo

import "github.com/orgming/mingdemo/framework"

type DemoProvider struct {
	framework.ServiceProvider

	c framework.Container
}

func (p *DemoProvider) Register(c framework.Container) framework.NewInstance {
	return NewService
}

func (p *DemoProvider) Boot(c framework.Container) error {
	p.c = c
	return nil
}

func (p *DemoProvider) IsDefer() bool {
	return false
}

func (p *DemoProvider) Params(c framework.Container) []any {
	return []any{p.c}
}

func (p *DemoProvider) Name() string {
	return DemoKey
}

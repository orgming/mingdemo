package redis

import (
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
)

type RedisProvider struct {
}

func (p *RedisProvider) Register(container framework.Container) framework.NewInstance {
	return NewMingRedis
}

func (p *RedisProvider) Boot(container framework.Container) error {
	return nil
}

func (p *RedisProvider) IsDefer() bool {
	return true
}

func (p *RedisProvider) Params(container framework.Container) []any {
	return []any{container}
}

func (p *RedisProvider) Name() string {
	return contract.RedisKey
}

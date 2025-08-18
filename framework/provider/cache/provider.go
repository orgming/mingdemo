package cache

import (
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
	"github.com/orgming/ming/framework/provider/cache/services"
	"strings"
)

type MingCacheProvider struct {
	framework.ServiceProvider

	Driver string
}

func (p *MingCacheProvider) Register(c framework.Container) framework.NewInstance {
	if p.Driver == "" {
		c, err := c.Make(contract.ConfigKey)
		if err != nil {
			return services.NewMemoryCache
		}
		cs := c.(contract.Config)
		p.Driver = strings.ToLower(cs.GetString("cache.driver"))

	}

	switch p.Driver {
	case "redis":
		return services.NewRedisCache
	case "memory":
		return services.NewMemoryCache
	default:
		return services.NewMemoryCache
	}
}

func (p *MingCacheProvider) Boot(c framework.Container) error {
	return nil
}

func (p *MingCacheProvider) IsDefer() bool {
	return true
}

func (p *MingCacheProvider) Params(c framework.Container) []any {
	return []any{c}
}

func (p *MingCacheProvider) Name() string {
	return contract.CacheKey
}

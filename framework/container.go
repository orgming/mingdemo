package framework

import (
	"errors"
	"fmt"
	"sync"
)

// Container 是一个服务容器，提供绑定服务提供者和获取服务实例的功能
type Container interface {
	// Bind 绑定一个服务提供者，如果关键字凭证已经存在，会进行替换操作，返回error
	Bind(provider ServiceProvider) error

	// IsBind 检查关键字凭证是否已经绑定了一个服务提供者
	IsBind(key string) bool

	// Make 根据关键字凭证获取一个服务
	Make(key string) (any, error)

	// MustMake 根据关键字凭证获取一个服务，如果这个关键字凭证未绑定服务提供者，那么会 panic。
	// 所以在使用这个接口的时候请保证服务容器已经为这个关键字凭证绑定了服务提供者。
	MustMake(key string) any

	// MakeNew 根据关键字凭证获取一个服务，只是这个服务并不是单例模式的
	// 它是根据服务提供者注册的启动函数和传递的 params 参数实例化出来的
	// 这个函数在需要为不同参数启动不同实例的时候非常有用
	// 根据不同参数获取新的实例
	MakeNew(key string, params []any) (any, error)
}

type MingContainer struct {
	Container // 强制要求实现Container接口
	// providers 存储注册的服务提供者，key为字符串凭证
	providers map[string]ServiceProvider
	// instances 具体的服务实例
	instances map[string]any
	// lock 用于锁住对容器的变更操作
	lock sync.RWMutex
}

func NewMingContainer() *MingContainer {
	return &MingContainer{
		providers: make(map[string]ServiceProvider),
		instances: make(map[string]any),
		lock:      sync.RWMutex{},
	}
}

// PrintProviders 输出服务容器中注册的关键字
func (mc *MingContainer) PrintProviders() []string {
	res := []string{}
	for _, provider := range mc.providers {
		res = append(res, fmt.Sprintf(provider.Name()))
	}
	return res
}

func (mc *MingContainer) Bind(provider ServiceProvider) error {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	key := provider.Name()

	mc.providers[key] = provider

	if provider.IsDefer() == false {
		if err := provider.Boot(mc); err != nil {
			return err
		}
		// 实例化方法
		params := provider.Params(mc)
		method := provider.Register(mc)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		mc.instances[key] = instance
	}
	return nil
}

func (mc *MingContainer) IsBind(key string) bool {
	return mc.findServiceProvider(key) != nil
}

func (mc *MingContainer) Make(key string) (any, error) {
	return mc.make(key, nil, false)
}

func (mc *MingContainer) MustMake(key string) any {
	serv, err := mc.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

func (mc *MingContainer) MakeNew(key string, params []any) (any, error) {
	return mc.make(key, params, true)
}

// 真正的实例化一个服务
func (mc *MingContainer) make(key string, params []any, forceNew bool) (any, error) {
	mc.lock.RLock()
	defer mc.lock.RUnlock()
	// 查询是否已经注册了这个服务提供者，如果没有注册，则返回错误
	sp := mc.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}

	if forceNew {
		return mc.newInstance(sp, params)
	}

	// 如果不是强制新实例化，那么就查询一下实例是否已经存在，如果存在，就直接返回
	if instance, ok := mc.instances[key]; ok {
		return instance, nil
	}

	// 如果容器中还未实例化，那么就实例化一个新的
	ins, err := mc.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}

	mc.instances[key] = ins
	return ins, nil

}

func (mc *MingContainer) findServiceProvider(key string) ServiceProvider {
	mc.lock.Lock()
	defer mc.lock.RUnlock()
	if sp, ok := mc.providers[key]; ok {
		return sp
	}
	return nil
}

func (mc *MingContainer) newInstance(sp ServiceProvider, params []any) (any, error) {
	if err := sp.Boot(mc); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(mc)
	}
	method := sp.Register(mc)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}

// NameList 列出容器中所有服务提供者的字符串凭证
func (mc *MingContainer) NameList() []string {
	ret := []string{}
	for _, provider := range mc.providers {
		name := provider.Name()
		ret = append(ret, name)
	}
	return ret
}

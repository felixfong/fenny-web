package framework

import (
	"errors"
	"fmt"
	"sync"
)

type Container interface {
	Bind(provider ServiceProvider) error

	Make(key string) (interface{}, error)

	IsBind(key string) bool

	MustMake(key string) interface{}

	MakeNew(key string, params []interface{}) (interface{}, error)
}

type HadeContainer struct {
	Container
	providers map[string]ServiceProvider
	instances map[string]interface{}

	lock sync.RWMutex
}

func NewHadeContainer() *HadeContainer {
	return &HadeContainer{
		providers: map[string]ServiceProvider{},
		instances: map[string]interface{}{},
		lock:      sync.RWMutex{},
	}
}

// PrintProviders 输出服务容器中注册的关键字
func (hade *HadeContainer) PrintProviders() []string {
	var ret []string
	for _, provider := range hade.providers {
		name := provider.Name()
		line := fmt.Sprint(name)
		ret = append(ret, line)
	}
	return ret
}

// Bind 将服务容器和关键字做绑定
func (hade *HadeContainer) Bind(provider ServiceProvider) error {
	hade.lock.Lock()
	defer hade.lock.Unlock()
	key := provider.Name()
	hade.providers[key] = provider

	if provider.IsDefer() == false {
		if err := provider.Boot(hade); err != nil {
			return err
		}
		// 实例化方法
		params := provider.Params(hade)
		method := provider.Register(hade)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		hade.instances[key] = instance
	}

	return nil
}

func (hade *HadeContainer) IsBind(key string) bool {
	return hade.findServiceProvider(key) != nil
}

func (hade *HadeContainer) findServiceProvider(key string) ServiceProvider {
	hade.lock.RLock()
	defer hade.lock.RUnlock()
	if sp, ok := hade.providers[key]; ok {
		return sp
	}
	return nil
}

func (hade *HadeContainer) Make(key string) (interface{}, error) {
	return hade.make(key, nil, false)
}

func (hade *HadeContainer) MustMake(key string) interface{} {
	serv, err := hade.make(key, nil, false)
	if err != nil {
		panic(err)
	}
	return serv
}

func (hade *HadeContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return hade.make(key, params, true)
}

// 实例化服务
func (hade *HadeContainer) make(key string, params []interface{}, forceNew bool) (interface{}, error) {
	hade.lock.RLock()
	defer hade.lock.RUnlock()
	sp := hade.findServiceProvider(key)
	if sp == nil {
		return nil, errors.New("contract " + key + " have not register")
	}
	if forceNew {
		return hade.newInstance(sp, params)
	}
	if ins, ok := hade.instances[key]; ok {
		return ins, nil
	}
	inst, err := hade.newInstance(sp, nil)
	if err != nil {
		return nil, err
	}
	hade.instances[key] = inst
	return inst, nil
}

func (hade *HadeContainer) newInstance(sp ServiceProvider, params []interface{}) (interface{}, error) {
	// force new a
	if err := sp.Boot(hade); err != nil {
		return nil, err
	}
	if params == nil {
		params = sp.Params(hade)
	}
	method := sp.Register(hade)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return ins, err
}



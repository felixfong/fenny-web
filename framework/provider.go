package framework

type NewInstance func(...interface{}) (interface{}, error)

// ServiceProvider 定义一个服务提供者需要实现的接口
type ServiceProvider interface {
	//Name 代表了这个服务提供者的凭证
	Name() string

	Register(Container) NewInstance

	Params(Container) []interface{}

	IsDefer() bool

	Boot(Container) error
}

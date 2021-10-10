package framework

type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)

	Group(string) IGroup

	Use(middlewares ...ControllerHandler)
}

type Group struct {
	core *Core
	prefix string

	middlewares []ControllerHandler
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		prefix: prefix,
	}
}

func (group *Group) Get(uri string, handler ControllerHandler) {
	uri = group.prefix + uri
	group.core.Get(uri, handler)
}

func (group *Group) Group(uri string) *Group{
	return NewGroup(group.core, uri)
}

func (group *Group) Use(middlewares ...ControllerHandler) {
	group.middlewares = middlewares
}
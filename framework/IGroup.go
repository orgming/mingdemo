package framework

// IGroup 代表前缀分组
type IGroup interface {
	Get(string, ControllerHandler)
	Post(string, ControllerHandler)
	Put(string, ControllerHandler)
	Delete(string, ControllerHandler)
}

type Group struct {
	core   *Core
	prefix string
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		prefix: prefix,
	}
}

func (g *Group) Get(url string, handler ControllerHandler) {
	g.core.Get(g.prefix+url, handler)
}

func (g *Group) Post(url string, handler ControllerHandler) {
	g.core.Post(g.prefix+url, handler)
}

func (g *Group) Put(url string, handler ControllerHandler) {
	g.core.Put(g.prefix+url, handler)
}

func (g *Group) Delete(url string, handler ControllerHandler) {
	g.core.Delete(g.prefix+url, handler)
}

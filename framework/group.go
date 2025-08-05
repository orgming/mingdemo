package framework

// IGroup 代表前缀分组
type IGroup interface {
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)

	Group(string) IGroup

	Use(middlewares ...ControllerHandler)
}

type Group struct {
	core   *Core
	prefix string
	parent *Group // 指向上一个Group，如果有的话

	middlewares []ControllerHandler
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:        core,
		prefix:      prefix,
		middlewares: []ControllerHandler{},
		parent:      nil,
	}
}

// Use 注册中间件
func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *Group) Get(url string, handlers ...ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	all := append(g.getMiddlewares(), handlers...)
	g.core.Get(url, all...)
}

func (g *Group) Post(url string, handlers ...ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	all := append(g.getMiddlewares(), handlers...)
	g.core.Post(url, all...)
}

func (g *Group) Put(url string, handlers ...ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	all := append(g.getMiddlewares(), handlers...)
	g.core.Put(url, all...)
}

func (g *Group) Delete(url string, handlers ...ControllerHandler) {
	url = g.getAbsolutePrefix() + url
	all := append(g.getMiddlewares(), handlers...)
	g.core.Delete(url, all...)
}

func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}

// 获取当前group的绝对路径
func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}

func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}
	return append(g.parent.getMiddlewares(), g.middlewares...)
}

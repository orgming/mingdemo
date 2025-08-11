package cobra

import "github.com/orgming/ming/framework"

func (c *Command) GetContainer() framework.Container {
	return c.container
}

func (c *Command) SetContainer(container framework.Container) {
	c.container = container
}

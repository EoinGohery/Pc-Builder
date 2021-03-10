package factory

type iCPU interface {
	setName(name string)
	getName() string
}

type cpu struct {
	name string
}

func (c *cpu) setName(name string) {
	c.name = name
}

func (c *cpu) getName() string {
	return c.name
}

package factory

import "fmt"

type iCPU interface {
	setName(name string)
	getName() string
	toString()
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

func (c *cpu) toString() {
	fmt.Printf(c.name)
}

package factory

import "fmt"

type iGPU interface {
	setName(name string)
	getName() string
	toString()
}

type gpu struct {
	name string
}

func (g *gpu) setName(name string) {
	g.name = name
}

func (g *gpu) getName() string {
	return g.name
}

func (g *gpu) toString() {
	fmt.Printf(g.name)
}

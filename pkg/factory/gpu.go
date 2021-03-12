package factory

type iGPU interface {
	setName(name string)
	getName() string
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

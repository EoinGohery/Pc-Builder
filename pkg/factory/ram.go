package factory

type iRAM interface {
	setName(name string)
	getName() string
}

type ram struct {
	name string
}

func (r *ram) setName(name string) {
	r.name = name
}

func (r *ram) getName() string {
	return r.name
}

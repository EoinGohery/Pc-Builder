package factory

import "fmt"

type iRAM interface {
	setName(name string)
	getName() string
	toString()
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

func (r *ram) toString() {
	fmt.Printf(r.name)
}

package factory

import "fmt"

type iDRIVE interface {
	setName(name string)
	getName() string
	toString()
}

type drive struct {
	name string
}

func (d *drive) setName(name string) {
	d.name = name
}

func (d *drive) getName() string {
	return d.name
}

func (d *drive) toString() {
	fmt.Printf(d.name)
}

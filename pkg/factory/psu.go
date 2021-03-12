package factory

import "fmt"

type iPSU interface {
	setName(name string)
	getName() string
	toString()
}

type psu struct {
	name string
}

func (p *psu) setName(name string) {
	p.name = name
}

func (p *psu) getName() string {
	return p.name
}

func (p *psu) toString() {
	fmt.Printf(p.name)
}

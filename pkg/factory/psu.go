package factory

type iPSU interface {
	setName(name string)
	getName() string
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

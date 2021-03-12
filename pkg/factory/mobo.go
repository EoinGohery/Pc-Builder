package factory

import "fmt"

type iMOBO interface {
	setName(name string)
	getName() string
	toString()
}

type mobo struct {
	name string
}

func (m *mobo) setName(name string) {
	m.name = name
}

func (m *mobo) getName() string {
	return m.name
}

func (m *mobo) toString() {
	fmt.Printf("Name: %s\n", m.name)
}

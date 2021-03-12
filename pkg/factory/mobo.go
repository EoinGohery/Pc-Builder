package factory

type iMOBO interface {
	setName(name string)
	getName() string
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

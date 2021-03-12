package factory

type iDRIVE interface {
	setName(name string)
	getName() string
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

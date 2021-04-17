package factory

type Component interface {
	ToString() string
	GetFilter() string
	Print()
}

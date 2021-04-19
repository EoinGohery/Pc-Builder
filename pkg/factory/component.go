package factory

type Component interface {
	ToString() string
	GetFilter() string
	Print()
	PrintIDString() string
	GetID() int
	clone() Component
}

package factory

type Component interface {
	ToString() string
	GetFilter() string
	Print()
	PrintIDString() string
	GetID() int
	SetID(id int)
	GetRamSlots() int
	GetDriveSlots() int
	clone() Component
	Add(Component)
}

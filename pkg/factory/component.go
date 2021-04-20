package factory

//component interfce to allow access to all part types from a slice od a single type.
//allows for composite structure of the pc
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

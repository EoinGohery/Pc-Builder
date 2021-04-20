package factory

import "fmt"

type iRAM interface {
	GetID() int
	GetPrice() int
	GetManufacturer() string
	GetName() string
	GetClock() string
	GetMemory() int
	SetData(id int, name string, price int, memory int, clock string)
	ToString() string
	PrintIDString() string
	Print()
	clone() Component
	GetFilter() string
	GetRamSlots() int
	GetDriveSlots() int
	SetID(id int)
	Add(Component)
}

type ram struct {
	id           int
	name         string
	manufacturer string
	price        int
	clock        string
	memory       int
}

// GetID returns the string id
func (r ram) GetID() int {
	return r.id
}

// SetID for int id
func (r *ram) SetID(id int) {
	r.id = id
}

// GetName returns the string name
func (r ram) GetName() string {
	return r.name
}

// GetManufacturer returns the string manufacturer
func (r ram) GetManufacturer() string {
	return r.manufacturer
}

// GetPrice returns the float32 price
func (r ram) GetPrice() int {
	return r.price
}

// GetClock returns the string clock
func (r ram) GetClock() string {
	return r.clock
}

// GetMemory returns the int memory
func (r ram) GetMemory() int {
	return r.memory
}

// GetFilter returns the filterable value if any
func (r ram) GetFilter() string {
	return ""
}

// GetDriveSlots returns the string id
func (r ram) GetDriveSlots() int {
	return 0
}

// GetRamSlots returns the string id
func (r ram) GetRamSlots() int {
	return 0
}

// Add does nothing for this part
func (r *ram) Add(Component) {

}

//set the data of a created object
func (r *ram) SetData(id int, name string, price int, memory int, clock string) {
	r.id = id
	r.name = name
	r.price = price
	r.memory = memory
	r.clock = clock
}

//prints the data with id
func (r *ram) Print() {
	fmt.Print(r.PrintIDString())
}

//string for final print statement to file
func (r *ram) ToString() string {
	return fmt.Sprintf("\nRam: %s %s Memory: %d Clock Speed: %s Price: %d", r.manufacturer, r.name, r.memory, r.clock, r.price)
}

//print statement for get all list
func (r *ram) PrintIDString() string {
	return fmt.Sprintf("\nID: %d Ram: %s %s Memory: %d Clock Speed: %s Price: %d", r.id, r.manufacturer, r.name, r.memory, r.clock, r.price)
}

//clone function to allow for prototyping
func (r *ram) clone() Component {
	return &ram{
		id:     r.id,
		name:   r.name + "_clone",
		price:  r.price,
		memory: r.memory,
		clock:  r.clock,
	}
}

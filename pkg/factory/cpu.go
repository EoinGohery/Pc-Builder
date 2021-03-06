package factory

import "fmt"

type iCPU interface {
	GetID() int
	GetPrice() int
	GetManufacturer() string
	GetCores() int
	GetClock() string
	GetSocket() string
	GetTpd() int
	SetData(id int, name string, price int, cores int, clock string, socket string, tdp int)
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

type cpu struct {
	id           int
	manufacturer string
	name         string
	price        int
	cores        int
	clock        string
	socket       string
	tdp          int
}

// GetID returns the string id
func (c cpu) GetID() int {
	return c.id
}

// SetID for int id
func (c *cpu) SetID(id int) {
	c.id = id
}

// GetManufacturer returns the string manufacturer
func (c cpu) GetManufacturer() string {
	return c.manufacturer
}

// GetName returns the string manufacturer
func (c cpu) GetName() string {
	return c.name
}

// GetPrice returns the float32 price
func (c cpu) GetPrice() int {
	return c.price
}

// GetCore returns the int cores
func (c cpu) GetCores() int {
	return c.cores
}

// GetClock returns the string clock
func (c cpu) GetClock() string {
	return c.clock
}

// GetSocket returns the string socket
func (c cpu) GetSocket() string {
	return c.socket
}

// GetTpd returns the int tdp
func (c cpu) GetTpd() int {
	return c.tdp
}

// GetFilter returns the filterable value if any
func (c cpu) GetFilter() string {
	return c.socket
}

// GetDriveSlots returns the string id
func (c cpu) GetDriveSlots() int {
	return 0
}

// GetRamSlots returns the string id
func (c cpu) GetRamSlots() int {
	return 0
}

// Add does nothing for this part
func (c *cpu) Add(Component) {

}

//set the data of a created object
func (c *cpu) SetData(id int, name string, price int, cores int, clock string, socket string, tdp int) {
	c.id = id
	c.name = name
	c.price = price
	c.cores = cores
	c.clock = clock
	c.socket = socket
	c.tdp = tdp
}

//prints the data with id
func (c *cpu) Print() {
	fmt.Print(c.PrintIDString())
}

//string for final print statement to file
func (c *cpu) ToString() string {
	return fmt.Sprintf("\nCPU: %s %s Cores: %d Clock Speed: %s Socket: %s tdp: %d Price: %d", c.manufacturer, c.name, c.cores, c.clock, c.socket, c.tdp, c.price)
}

//print statement for get all list
func (c *cpu) PrintIDString() string {
	return fmt.Sprintf("\nID: %v CPU: %s %s Cores: %d Clock Speed: %s Socket: %s tdp: %d Price: %d", c.id, c.manufacturer, c.name, c.cores, c.clock, c.socket, c.tdp, c.price)
}

//clone function to allow for prototyping
func (c *cpu) clone() Component {
	return &cpu{
		id:     c.id,
		name:   c.name + "_clone",
		price:  c.price,
		cores:  c.cores,
		clock:  c.clock,
		socket: c.socket,
		tdp:    c.tdp,
	}
}

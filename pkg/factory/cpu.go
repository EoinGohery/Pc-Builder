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

func (c *cpu) SetData(id int, name string, price int, cores int, clock string, socket string, tdp int) {
	c.id = id
	c.name = name
	c.price = price
	c.cores = cores
	c.clock = clock
	c.socket = socket
	c.tdp = tdp
}

func (c *cpu) Print() {
	fmt.Print(c.PrintIDString())
}

func (c *cpu) ToString() string {
	return fmt.Sprintf("\n CPU: %s %s Cores: %d Clock Speed: %s Socket: %s tdp: %d Price: %d", c.manufacturer, c.name, c.cores, c.clock, c.socket, c.tdp, c.price)
}

func (c *cpu) PrintIDString() string {
	return fmt.Sprintf("\nID: %v CPU: %s %s Cores: %d Clock Speed: %s Socket: %s tdp: %d Price: %d", c.id, c.manufacturer, c.name, c.cores, c.clock, c.socket, c.tdp, c.price)
}

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

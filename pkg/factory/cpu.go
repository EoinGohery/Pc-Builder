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
	SetData(id int, name string, price int, cores int, clock string, socket string, tpd int)
	ToString() string
	Print()
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
	tpd          int
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

// GetTpd returns the int tpd
func (c cpu) GetTpd() int {
	return c.tpd
}

// GetFilter returns the filterable value if any
func (c cpu) GetFilter() string {
	return c.socket
}

func (c *cpu) SetData(id int, name string, price int, cores int, clock string, socket string, tpd int) {
	c.id = id
	c.name = name
	c.price = price
	c.cores = cores
	c.clock = clock
	c.socket = socket
	c.tpd = tpd
}

func (c *cpu) Print() {
	fmt.Print(c.ToString())
}

func (c *cpu) ToString() string {
	return fmt.Sprintf("\nCPU: %s %s Cores: %d Clock Speed: %s Socket: %s TPD: %d Price: %d", c.manufacturer, c.name, c.cores, c.clock, c.socket, c.tpd, c.price)
}

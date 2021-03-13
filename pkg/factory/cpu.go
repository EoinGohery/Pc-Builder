package factory

import "fmt"

type iCPU interface {
	GetPrice() float32
	GetManufacturer() string
	GetCores() int
	GetClock() string
	GetSocket() string
	GetTpd() int
	setData(name string, price float32, cores int, clock string, socket string, tpd int)
	toString() string
	print()
}

type cpu struct {
	manufacturer string
	name         string
	price        float32
	cores        int
	clock        string
	socket       string
	tpd          int
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
func (c cpu) GetPrice() float32 {
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

func (c *cpu) setData(name string, price float32, cores int, clock string, socket string, tpd int) {
	c.name = name
	c.price = price
	c.cores = cores
	c.clock = clock
	c.socket = socket
	c.tpd = tpd
}

func (c *cpu) print() {
	fmt.Printf(c.toString())
}

func (c *cpu) toString() string {
	return fmt.Sprintf("\nCPU: %s %s Cores: %d Clock Speed: %s Socket: %s TPD: %d Price: %f", c.manufacturer, c.name, c.cores, c.clock, c.socket, c.tpd, c.price)
}

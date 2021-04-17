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
	Print()
	GetFilter() string
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

func (r *ram) SetData(id int, name string, price int, memory int, clock string) {
	r.id = id
	r.name = name
	r.price = price
	r.memory = memory
	r.clock = clock
}

func (r *ram) Print() {
	fmt.Print(r.ToString())
}

func (r *ram) ToString() string {
	return fmt.Sprintf("\nRam: %s %s Memory: %d Clock Speed: %s Price: %d", r.manufacturer, r.name, r.memory, r.clock, r.price)
}

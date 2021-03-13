package factory

import "fmt"

type iRAM interface {
	GetPrice() float32
	GetManufacturer() string
	GetName() string
	GetClock() string
	GetMemory() int
	setData(name string, price float32, memory int, clock string)
	toString() string
	print()
}

type ram struct {
	name         string
	manufacturer string
	price        float32
	clock        string
	memory       int
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
func (r ram) GetPrice() float32 {
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

func (r *ram) setData(name string, price float32, memory int, clock string) {
	r.name = name
	r.price = price
	r.memory = memory
	r.clock = clock
}

func (r *ram) print() {
	fmt.Printf(r.toString())
}

func (r *ram) toString() string {
	return fmt.Sprintf("\nRam: %s %s Memory: %d Clock Speed: %s Price: %f", r.manufacturer, r.name, r.memory, r.clock, r.price)
}

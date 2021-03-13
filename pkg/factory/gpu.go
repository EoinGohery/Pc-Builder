package factory

import "fmt"

type iGPU interface {
	GetPrice() float32
	GetName() string
	GetManufacturer() string
	GetMemory() int
	GetClock() string
	setData(name string, price float32, tpd int, memory int, clock string)
	toString() string
	print()
}

type gpu struct {
	name         string
	manufacturer string
	price        float32
	memory       int
	clock        string
	tpd          int
}

// GetName returns the string name
func (g gpu) GetName() string {
	return g.name
}

// GetManufacturer returns the string manufacturer
func (g gpu) GetManufacturer() string {
	return g.manufacturer
}

// GetPrice returns the float32 price
func (g gpu) GetPrice() float32 {
	return g.price
}

// GetMemory returns the int memory
func (g gpu) GetMemory() int {
	return g.memory
}

// GetClock returns the string clock
func (g gpu) GetClock() string {
	return g.clock
}

// GetTpd returns the int tpd
func (g gpu) GetTpd() int {
	return g.tpd
}

func (g *gpu) setData(name string, price float32, tpd int, memory int, clock string) {
	g.name = name
	g.price = price
	g.tpd = tpd
	g.memory = memory
	g.clock = clock
}

func (g *gpu) print() {
	fmt.Printf(g.toString())
}

func (g *gpu) toString() string {
	return fmt.Sprintf("\nGPU: %s %s Clock Speed: %s Memory: %d TPD: %d Price: %f", g.manufacturer, g.name, g.clock, g.memory, g.tpd, g.price)
}

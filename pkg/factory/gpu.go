package factory

import "fmt"

type iGPU interface {
	GetID() int
	GetPrice() int
	GetName() string
	GetManufacturer() string
	GetMemory() int
	GetClock() string
	SetData(id int, name string, price int, tpd int, memory int, clock string)
	ToString() string
	Print()
	GetFilter() string
}

type gpu struct {
	id           int
	name         string
	manufacturer string
	price        int
	memory       int
	clock        string
	tpd          int
}

// GetID returns the string id
func (g gpu) GetID() int {
	return g.id
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
func (g gpu) GetPrice() int {
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

// GetFilter returns the filterable value if any
func (g gpu) GetFilter() string {
	return ""
}

func (g *gpu) SetData(id int, name string, price int, tpd int, memory int, clock string) {
	g.id = id
	g.name = name
	g.price = price
	g.tpd = tpd
	g.memory = memory
	g.clock = clock
}

func (g *gpu) Print() {
	fmt.Print(g.ToString())
}

func (g *gpu) ToString() string {
	return fmt.Sprintf("\nGPU: %s %s Clock Speed: %s Memory: %d TPD: %d Price: %d", g.manufacturer, g.name, g.clock, g.memory, g.tpd, g.price)
}

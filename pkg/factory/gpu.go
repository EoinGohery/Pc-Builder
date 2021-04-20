package factory

import "fmt"

type iGPU interface {
	GetID() int
	GetPrice() int
	GetName() string
	GetManufacturer() string
	GetMemory() int
	GetClock() string
	SetData(id int, name string, price int, tdp int, memory int, clock string)
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

type gpu struct {
	id           int
	name         string
	manufacturer string
	price        int
	memory       int
	clock        string
	tdp          int
}

// GetID returns the string id
func (g gpu) GetID() int {
	return g.id
}

// SetID for int id
func (g *gpu) SetID(id int) {
	g.id = id
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

// GetTpd returns the int tdp
func (g gpu) GetTpd() int {
	return g.tdp
}

// GetFilter returns the filterable value if any
func (g gpu) GetFilter() string {
	return ""
}

// GetDriveSlots returns the string id
func (g gpu) GetDriveSlots() int {
	return 0
}

// GetRamSlots returns the string id
func (g gpu) GetRamSlots() int {
	return 0
}

// Add does nothing for this part
func (g *gpu) Add(Component) {

}

func (g *gpu) SetData(id int, name string, price int, tdp int, memory int, clock string) {
	g.id = id
	g.name = name
	g.price = price
	g.tdp = tdp
	g.memory = memory
	g.clock = clock
}

func (g *gpu) Print() {
	fmt.Print(g.PrintIDString())
}

func (g *gpu) ToString() string {
	return fmt.Sprintf("\nGPU: %s %s  Clock Speed: %s Memory: %d tdp: %d Price: %d", g.manufacturer, g.name, g.clock, g.memory, g.tdp, g.price)
}

func (g *gpu) PrintIDString() string {
	return fmt.Sprintf("\nID: %v GPU: %s %s  Clock Speed: %s Memory: %d tdp: %d Price: %d", g.id, g.manufacturer, g.name, g.clock, g.memory, g.tdp, g.price)
}

func (g *gpu) clone() Component {
	return &gpu{
		id:     g.id,
		name:   g.name + "_clone",
		price:  g.price,
		tdp:    g.tdp,
		memory: g.memory,
		clock:  g.clock,
	}
}

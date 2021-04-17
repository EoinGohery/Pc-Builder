package factory

import "fmt"

type iMOBO interface {
	GetID() int
	GetManufacturer() string
	GetName() string
	GetPrice() int
	GetSocket() string
	GetTpd() int
	GetRamSlots() int
	GetMaxRam() int
	GetDriveSlots() int
	SetData(id int, name string, price int, ramSlots int, maxRam int, driveSlots int, socket string, tdp int)
	ToString() string
	Print()
	GetFilter() string
	Add(c Component)
}

type mobo struct {
	components   []Component
	id           int
	manufacturer string
	name         string
	price        int
	socket       string
	tdp          int
	ramSlots     int
	maxRam       int
	driveSlots   int
}

func (m *mobo) Add(c Component) {
	m.components = append(m.components, c)
}

// GetID returns the string id
func (m mobo) GetID() int {
	return m.id
}

// GetManufacturer returns the string manufacturer
func (m mobo) GetManufacturer() string {
	return m.manufacturer
}

// GetName returns the string name
func (m mobo) GetName() string {
	return m.name
}

// GetPrice returns the float32 price
func (m mobo) GetPrice() int {
	return m.price
}

// GetSocket returns the string socket
func (m mobo) GetSocket() string {
	return m.socket
}

// GetTpd returns the int tdp
func (m mobo) GetTpd() int {
	return m.tdp
}

// GetRamSlot returns the int ramSlots
func (m mobo) GetRamSlots() int {
	return m.ramSlots
}

// GetMaxRam returns the int maxRam
func (m mobo) GetMaxRam() int {
	return m.maxRam
}

// GetDriveSlot returns the int driveSlots
func (m mobo) GetDriveSlots() int {
	return m.driveSlots
}

// GetFilter returns the filterable value if any
func (m mobo) GetFilter() string {
	return m.socket
}

func (m *mobo) SetData(id int, name string, price int, ramSlots int, maxRam int, driveSlots int, socket string, tdp int) {
	m.id = id
	m.name = name
	m.price = price
	m.socket = socket
	m.tdp = tdp
	m.ramSlots = ramSlots
	m.maxRam = maxRam
	m.driveSlots = driveSlots
}

func (m *mobo) Print() {
	fmt.Print(m.ToString())
	for _, composite := range m.components {
		composite.Print()
	}
}

func (m *mobo) ToString() string {
	return fmt.Sprintf("\nMotherboard: %s %s Max Ram: %d Ram Slots: %d  Drive Slots: %d TPD: %d Price: %d", m.manufacturer, m.name, m.maxRam, m.ramSlots, m.driveSlots, m.tdp, m.price)
}

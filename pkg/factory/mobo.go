package factory

import "fmt"

type iMOBO interface {
	GetManufacturer() string
	GetName() string
	GetPrice() float32
	GetSocket() string
	GetTpd() int
	GetRamSlots() int
	GetMaxRam() int
	GetDriveSlots() int
	setData(name string, price float32, ramSlots int, maxRam int, driveSlots int, socket string, tpd int)
	toString() string
	print()
	add(c component)
}

type mobo struct {
	components   []component
	manufacturer string
	name         string
	price        float32
	socket       string
	tpd          int
	ramSlots     int
	maxRam       int
	driveSlots   int
}

func (m *mobo) add(c component) {
	m.components = append(m.components, c)
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
func (m mobo) GetPrice() float32 {
	return m.price
}

// GetSocket returns the string socket
func (m mobo) GetSocket() string {
	return m.socket
}

// GetTpd returns the int tpd
func (m mobo) GetTpd() int {
	return m.tpd
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

func (m *mobo) setData(name string, price float32, ramSlots int, maxRam int, driveSlots int, socket string, tpd int) {
	m.name = name
	m.price = price
	m.socket = socket
	m.tpd = tpd
	m.ramSlots = ramSlots
	m.maxRam = maxRam
	m.driveSlots = driveSlots
}

func (m *mobo) print() {
	fmt.Printf(m.toString())
	for _, composite := range m.components {
		composite.print()
	}
}

func (m *mobo) toString() string {
	return fmt.Sprintf("\nMotherboard: %s %s Max Ram: %d Ram Slots: %d  Drive Slots: %d TPD: %d Price: %f", m.manufacturer, m.name, m.maxRam, m.ramSlots, m.driveSlots, m.tpd, m.price)
}

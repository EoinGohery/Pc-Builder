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
	PrintIDString() string
	Print()
	clone() Component
	GetFilter() string
	Add(c Component)
	SetID(id int)
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

// SetID for int id
func (m *mobo) SetID(id int) {
	m.id = id
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
	fmt.Print(m.PrintIDString())
	for _, composite := range m.components {
		composite.Print()
	}
}

func (m *mobo) ToString() string {
	var result = fmt.Sprintf("\nMotherboard: %s %s Socket %s Max Ram: %d Ram Slots: %d  Drive Slots: %d TPD: %d Price: %d", m.manufacturer, m.name, m.socket, m.maxRam, m.ramSlots, m.driveSlots, m.tdp, m.price)
	for _, composite := range m.components {
		result = fmt.Sprintf("%s%s", result, composite.ToString())
	}
	return result
}

func (m *mobo) PrintIDString() string {
	return fmt.Sprintf("\nID: %d Motherboard: %s %s Socket %s Max Ram: %d Ram Slots: %d  Drive Slots: %d TPD: %d Price: %d", m.id, m.manufacturer, m.name, m.socket, m.maxRam, m.ramSlots, m.driveSlots, m.tdp, m.price)
}

func (m *mobo) clone() Component {
	cloneBuild := &mobo{
		id:         m.id,
		name:       m.name + "_clone",
		price:      m.price,
		socket:     m.socket,
		tdp:        m.tdp,
		ramSlots:   m.ramSlots,
		maxRam:     m.maxRam,
		driveSlots: m.driveSlots,
	}
	var tempComponent []Component
	for _, i := range m.components {
		copy := i.clone()
		tempComponent = append(tempComponent, copy)
	}
	cloneBuild.components = tempComponent
	return cloneBuild
}

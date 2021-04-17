package factory

import "fmt"

type iDRIVE interface {
	GetID() int
	GetPrice() int
	GetManufacturer() string
	GetName() string
	GetSize() int
	SetData(id int, name string, price int, tpd int, size int, technology string)
	ToString() string
	Print()
	GetFilter() string
}

type drive struct {
	id           int
	name         string
	manufacturer string
	price        int
	size         int
	tpd          int
	technology   string
}

// GetID returns the string id
func (d drive) GetID() int {
	return d.id
}

// GetTpd returns the int tpd
func (d drive) GetTpd() int {
	return d.tpd
}

// GetName returns the string name
func (d drive) GetName() string {
	return d.name
}

// GetManufacturer returns the string manufacturer
func (d drive) GetManufacturer() string {
	return d.manufacturer
}

// GetPrice returns the float32 price
func (d drive) GetPrice() int {
	return d.price
}

// GetSize returns the int size
func (d drive) GetSize() int {
	return d.size
}

// GetTechnology returns the string technology
func (d drive) GetTechnology() string {
	return d.technology
}

// GetFilter returns the filterable value if any
func (d drive) GetFilter() string {
	return ""
}

func (d *drive) SetData(id int, name string, price int, tpd int, size int, technology string) {
	d.id = id
	d.name = name
	d.price = price
	d.tpd = tpd
	d.size = size
	d.technology = technology
}

func (d *drive) Print() {
	fmt.Print(d.ToString())
}

func (d *drive) ToString() string {
	return fmt.Sprintf("\nDrive: %s %s %s Size: %d TPD: %d Price: %d", d.manufacturer, d.name, d.technology, d.size, d.tpd, d.price)
}

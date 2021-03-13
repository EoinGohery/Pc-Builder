package factory

import "fmt"

type iDRIVE interface {
	GetPrice() float32
	GetManufacturer() string
	GetName() string
	GetSize() int
	setData(name string, price float32, tpd int, size int, technology string)
	toString() string
	print()
}

type drive struct {
	name         string
	manufacturer string
	price        float32
	size         int
	tpd          int
	technology   string
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
func (d drive) GetPrice() float32 {
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

func (d *drive) setData(name string, price float32, tpd int, size int, technology string) {
	d.name = name
	d.price = price
	d.tpd = tpd
	d.size = size
	d.technology = technology
}

func (d *drive) print() {
	fmt.Printf(d.toString())
}

func (d *drive) toString() string {
	return fmt.Sprintf("\nDrive: %s %s %s Size: %d TPD: %d Price: %f", d.manufacturer, d.name, d.technology, d.size, d.tpd, d.price)
}

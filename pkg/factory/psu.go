package factory

import "fmt"

type iPSU interface {
	GetPrice() float32
	GetManufacturer() string
	GetName() string
	GetRating() string
	GetCapacity() int
	setData(name string, price float32, capacity int, rating string)
	toString() string
	print()
}

type psu struct {
	name         string
	manufacturer string
	price        float32
	rating       string
	capacity     int
}

// GetName returns the string name
func (p psu) GetName() string {
	return p.name
}

// GetManufacturer returns the string manufacturer
func (p psu) GetManufacturer() string {
	return p.manufacturer
}

// GetPrice returns the float32 price
func (p psu) GetPrice() float32 {
	return p.price
}

// GetRating returns the string rating
func (p psu) GetRating() string {
	return p.rating
}

// GetCapacity returns the int capacity
func (p psu) GetCapacity() int {
	return p.capacity
}

func (p *psu) setData(name string, price float32, capacity int, rating string) {
	p.name = name
	p.price = price
	p.capacity = capacity
	p.rating = rating
}

func (p *psu) print() {
	fmt.Printf(p.toString())
}

func (p *psu) toString() string {
	return fmt.Sprintf("\nMotherboard: %s %s Output: %d W Rating: %s Price: %f", p.manufacturer, p.name, p.capacity, p.rating, p.price)
}

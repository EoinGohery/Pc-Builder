package factory

import "fmt"

type iPSU interface {
	GetID() int
	GetPrice() int
	GetManufacturer() string
	GetName() string
	GetRating() string
	GetCapacity() int
	SetData(id int, name string, price int, capacity int, rating string)
	ToString() string
	PrintIDString() string
	Print()
	clone() Component
	GetFilter() string
}

type psu struct {
	id           int
	name         string
	manufacturer string
	price        int
	rating       string
	capacity     int
}

// GetID returns the string id
func (p psu) GetID() int {
	return p.id
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
func (p psu) GetPrice() int {
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

// GetFilter returns the filterable value if any
func (p psu) GetFilter() string {
	return ""
}

func (p *psu) SetData(id int, name string, price int, capacity int, rating string) {
	p.id = id
	p.name = name
	p.price = price
	p.capacity = capacity
	p.rating = rating
}

func (p *psu) Print() {
	fmt.Print(p.PrintIDString())
}

func (p *psu) ToString() string {
	return fmt.Sprintf("\nMotherboard: %s %s Output: %d W Rating: %s Price: %d", p.manufacturer, p.name, p.capacity, p.rating, p.price)
}

func (p *psu) PrintIDString() string {
	return fmt.Sprintf("\nID: %d Motherboard: %s %s Output: %d W Rating: %s Price: %d", p.id, p.manufacturer, p.name, p.capacity, p.rating, p.price)
}

func (p *psu) clone() Component {
	return &psu{
		id:       p.id,
		name:     p.name + "_clone",
		price:    p.price,
		capacity: p.capacity,
		rating:   p.rating,
	}
}

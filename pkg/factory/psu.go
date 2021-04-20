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
	GetRamSlots() int
	GetDriveSlots() int
	SetID(id int)
	Add(Component)
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

// SetID for int id
func (p *psu) SetID(id int) {
	p.id = id
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

// GetDriveSlots returns the string id
func (p psu) GetDriveSlots() int {
	return 0
}

// GetRamSlots returns the string id
func (p psu) GetRamSlots() int {
	return 0
}

// Add does nothing for this part
func (p *psu) Add(Component) {

}

//set the data of a created object
func (p *psu) SetData(id int, name string, price int, capacity int, rating string) {
	p.id = id
	p.name = name
	p.price = price
	p.capacity = capacity
	p.rating = rating
}

//prints the data with id
func (p *psu) Print() {
	fmt.Print(p.PrintIDString())
}

//string for final print statement to file
func (p *psu) ToString() string {
	return fmt.Sprintf("\nPSU: %s %s Output: %d W Rating: %s Price: %d", p.manufacturer, p.name, p.capacity, p.rating, p.price)

}

//print statement for get all list
func (p *psu) PrintIDString() string {
	return fmt.Sprintf("\nID: %d PSU: %s %s Output: %d W Rating: %s Price: %d", p.id, p.manufacturer, p.name, p.capacity, p.rating, p.price)
}

//clone function to allow for prototyping
func (p *psu) clone() Component {
	return &psu{
		id:       p.id,
		name:     p.name + "_clone",
		price:    p.price,
		capacity: p.capacity,
		rating:   p.rating,
	}
}

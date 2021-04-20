package factory

import "fmt"

type Tower struct {
	components []Component
	name       string
}

//used to add the componenes to teh composite structure
func (t *Tower) Add(c Component) {

	t.components = append(t.components, c)
}

//prints the data with id
func (t *Tower) Print() {
	fmt.Printf("\nPrinting parts list for selected build:")
	for _, composite := range t.components {
		composite.Print()
	}
}

//recuressuvly loops through all componments and returns a single string of all data
func (t *Tower) ToString() string {
	var result string
	for _, composite := range t.components {
		result = fmt.Sprintf("%s%s", result, composite.ToString())
	}
	return result
}

//clone function to allow for recurssive prototyping of components contained in teh object
func (t *Tower) Clone() Tower {
	cloneBuild := &Tower{name: t.name + "_clone"}
	var tempComponent []Component
	for _, i := range t.components {
		copy := i.clone()
		tempComponent = append(tempComponent, copy)
	}
	cloneBuild.components = tempComponent
	return *cloneBuild
}

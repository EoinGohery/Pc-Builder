package factory

import "fmt"

type Tower struct {
	components []Component
}

func (t *Tower) Add(c Component) {
	t.components = append(t.components, c)
}

func (t *Tower) Print() {
	fmt.Printf("\nPrinting parts list for selected build:")
	for _, composite := range t.components {
		composite.Print()
	}
}

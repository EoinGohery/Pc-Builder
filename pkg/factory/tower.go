package factory

import "fmt"

type tower struct {
	components []component
}

func (t *tower) add(c component) {
	t.components = append(t.components, c)
}

func (t *tower) Print() {
	fmt.Printf("\nPrinting parts list for selected build:")
	for _, composite := range t.components {
		composite.print()
	}
}

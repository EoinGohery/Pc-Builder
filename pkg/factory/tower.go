package factory

import "fmt"

type Tower struct {
	components []Component
	name       string
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

func (t *Tower) ToString() string {
	var result string
	for _, composite := range t.components {
		result = fmt.Sprintf("%s%s", result, composite.ToString())
	}
	return result
}

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

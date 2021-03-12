package factory

import "fmt"

//PrintDetails esport to display system info
func PrintDetails() {
	asusFactory, _ := getPartsFactory("asus")
	amdFactory, _ := getPartsFactory("nike")

	amdMOBO := asusFactory.makeMOBO()

	amdCPU := amdFactory.makeCPU()

	printCPUDetails(amdCPU)

	printMOBODetails(amdMOBO)
}

func printCPUDetails(g iCPU) {
	fmt.Printf("Logo: %s", g.getName())
	fmt.Println()
	fmt.Printf("Size: %s", g.getName())
	fmt.Println()
}

func printMOBODetails(m iMOBO) {
	fmt.Printf("Size: %s", m.getName())
	fmt.Println()
	fmt.Printf("Size: %s", m.getName())
	fmt.Println()

}

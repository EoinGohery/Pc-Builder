package factory

//PrintDetails esport to display system info
func PrintDetails() {

	tower := &tower{}

	asusFactory, _ := getPartsFactory("asus")
	amdFactory, _ := getPartsFactory("amd")
	seagateFactory, _ := getPartsFactory("seagate")
	nvidiaFactory, _ := getPartsFactory("nvidia")

	mobo := asusFactory.makeMOBO()
	//mobo.setData()
	tower.add(mobo)

	cpu := amdFactory.makeCPU()
	tower.add(cpu)

	ram := asusFactory.makeRAM()
	tower.add(ram)

	drive := seagateFactory.makeDRIVE()
	tower.add(drive)

	gpu := nvidiaFactory.makeGPU()
	tower.add(gpu)

	psu := asusFactory.makePSU()
	tower.add(psu)

	tower.Print()

}

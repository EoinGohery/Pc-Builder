package factory

//PrintDetails esport to display system info
func PrintDetails() {

	tower := &tower{}

	asusFactory, _ := getPartsFactory("asus")
	amdFactory, _ := getPartsFactory("amd")
	seagateFactory, _ := getPartsFactory("seagate")
	nvidiaFactory, _ := getPartsFactory("nvidia")

	mobo := asusFactory.makeMOBO()

	cpu := amdFactory.makeCPU()
	ram := asusFactory.makeRAM()
	tower.add(cpu)
	tower.add(ram)
	tower.add(mobo)

	drive := seagateFactory.makeDRIVE()
	tower.add(drive)

	gpu := nvidiaFactory.makeGPU()
	tower.add(gpu)

	psu := asusFactory.makePSU()
	tower.add(psu)

	tower.toString()

}

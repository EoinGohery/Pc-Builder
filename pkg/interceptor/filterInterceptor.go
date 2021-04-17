package interceptor

import (
	"CS4227/pkg/factory"
	"CS4227/pkg/filter"
)

type FilterInterceptor struct {
}

func (f FilterInterceptor) execute(request string, data string, filtertoSet string) []factory.Component {
	formatted := filter.JsonParse(request, data)
	filtered := filter.Filter(formatted, filtertoSet)
	return filtered
}

func PrintDetails() {

	tower := &factory.Tower{}

	asusFactory, _ := factory.GetPartsFactory("asus")
	amdFactory, _ := factory.GetPartsFactory("amd")
	seagateFactory, _ := factory.GetPartsFactory("seagate")
	nvidiaFactory, _ := factory.GetPartsFactory("nvidia")

	mobo := asusFactory.MakeMOBO()
	//mobo.setData()
	tower.Add(mobo)

	cpu := amdFactory.MakeCPU()
	tower.Add(cpu)

	ram := asusFactory.MakeRAM()
	tower.Add(ram)

	drive := seagateFactory.MakeDRIVE()
	tower.Add(drive)

	gpu := nvidiaFactory.MakeGPU()
	tower.Add(gpu)

	psu := asusFactory.MakePSU()
	tower.Add(psu)

	tower.Print()

}

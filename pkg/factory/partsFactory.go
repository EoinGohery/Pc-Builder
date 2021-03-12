package factory

import "fmt"

type partsFactory interface {
	makeCPU() iCPU
	makeGPU() iGPU
	makeMOBO() iMOBO
	makeRAM() iRAM
	makeDRIVE() iDRIVE
	makePSU() iPSU
}

//GetPartsFactory used to iniitalise parts
func getPartsFactory(brand string) (partsFactory, error) {
	if brand == "amd" {
		return &amd{}, nil
	} else if brand == "intel" {
		return &intel{}, nil
	} else if brand == "nvidia" {
		return &nvidia{}, nil
	} else if brand == "asus" {
		return &asus{}, nil
	} else if brand == "seagate" {
		return &seagate{}, nil
	}

	return nil, fmt.Errorf("Wrong Manufacturer passed")
}

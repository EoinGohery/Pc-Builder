package factory

import "fmt"

type factory interface {
	makeCPU() iCPU
	makeGPU() iGPU
	makeMOBO() iMOBO
	makeRAM() iRAM
	makeDRIVE() iDRIVE
	makePSU() iPSU
}

func getPartsFactory(brand string) (factory, error) {
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

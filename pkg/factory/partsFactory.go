package factory

import "fmt"

type partsFactory interface {
	MakeCPU() iCPU
	MakeGPU() iGPU
	MakeMOBO() iMOBO
	MakeRAM() iRAM
	MakeDRIVE() iDRIVE
	MakePSU() iPSU
}

//GetPartsFactory used to initalise parts
func GetPartsFactory(brand string) (partsFactory, error) {
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

	return nil, fmt.Errorf("wrong manufacturer passed")
}

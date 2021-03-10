package factory

import "fmt"

type partsFactory interface {
	makeCPU() iCPU
	makeGPU() iGPU
	makeMobo() iMOBO
	makeRam() iRam
	makeHardDrive() iDrive
	makePSU() iPSU
}

func getPartsFactory(brand string) (partsFactory, error) {
	if brand == "amd" {
		return &amd{}, nil
	}

	return nil, fmt.Errorf("Wrong Manufacturer passed")
}

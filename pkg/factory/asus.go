package factory

type asus struct {
}

func (a *asus) makeCPU() iCPU {
	return &asusCPU{
		cpu: cpu{
			manufacturer: "Asus",
		},
	}
}

func (a *asus) makeGPU() iGPU {
	return &asusGPU{
		gpu: gpu{
			manufacturer: "Asus",
		},
	}
}

func (a *asus) makeMOBO() iMOBO {
	return &asusMOBO{
		mobo: mobo{
			manufacturer: "Asus",
		},
	}
}

func (a *asus) makeRAM() iRAM {
	return &asusRAM{
		ram: ram{
			manufacturer: "Asus",
		},
	}
}

func (a *asus) makeDRIVE() iDRIVE {
	return &asusDRIVE{
		drive: drive{
			manufacturer: "Asus",
		},
	}
}

func (a *asus) makePSU() iPSU {
	return &asusPSU{
		psu: psu{
			manufacturer: "Asus",
		},
	}
}

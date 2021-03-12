package factory

type asus struct {
}

func (a *asus) makeCPU() iCPU {
	return &asusCPU{
		cpu: cpu{
			name: "Ryzen",
		},
	}
}

func (a *asus) makeGPU() iGPU {
	return &asusGPU{
		gpu: gpu{
			name: "Ryzen",
		},
	}
}

func (a *asus) makeMOBO() iMOBO {
	return &asusMOBO{
		mobo: mobo{
			name: "Ryzen",
		},
	}
}

func (a *asus) makeRAM() iRAM {
	return &asusRAM{
		ram: ram{
			name: "Ryzen",
		},
	}
}

func (a *asus) makeDRIVE() iDRIVE {
	return &asusDRIVE{
		drive: drive{
			name: "Ryzen",
		},
	}
}

func (a *asus) makePSU() iPSU {
	return &asusPSU{
		psu: psu{
			name: "Ryzen",
		},
	}
}

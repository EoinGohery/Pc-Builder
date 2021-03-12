package factory

type amd struct {
}

func (a *amd) makeCPU() iCPU {
	return &amdCPU{
		cpu: cpu{
			name: "Ryzen",
		},
	}
}

func (a *amd) makeGPU() iGPU {
	return &amdGPU{
		gpu: gpu{
			name: "Ryzen",
		},
	}
}

func (a *amd) makeMOBO() iMOBO {
	return &amdMOBO{
		mobo: mobo{
			name: "Ryzen",
		},
	}
}

func (a *amd) makeRAM() iRAM {
	return &amdRAM{
		ram: ram{
			name: "Ryzen",
		},
	}
}

func (a *amd) makeDRIVE() iDRIVE {
	return &amdDRIVE{
		drive: drive{
			name: "Ryzen",
		},
	}
}

func (a *amd) makePSU() iPSU {
	return &amdPSU{
		psu: psu{
			name: "Ryzen",
		},
	}
}

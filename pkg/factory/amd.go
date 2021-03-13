package factory

type amd struct {
}

func (a *amd) makeCPU() iCPU {
	return &amdCPU{
		cpu: cpu{
			manufacturer: "Amd",
		},
	}
}

func (a *amd) makeGPU() iGPU {
	return &amdGPU{
		gpu: gpu{
			manufacturer: "Amd",
		},
	}
}

func (a *amd) makeMOBO() iMOBO {
	return &amdMOBO{
		mobo: mobo{
			manufacturer: "Amd",
		},
	}
}

func (a *amd) makeRAM() iRAM {
	return &amdRAM{
		ram: ram{
			manufacturer: "Amd",
		},
	}
}

func (a *amd) makeDRIVE() iDRIVE {
	return &amdDRIVE{
		drive: drive{
			manufacturer: "Amd",
		},
	}
}

func (a *amd) makePSU() iPSU {
	return &amdPSU{
		psu: psu{
			manufacturer: "Amd",
		},
	}
}

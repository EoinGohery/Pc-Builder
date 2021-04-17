package factory

type amd struct {
}

func (a *amd) MakeCPU() iCPU {
	return &amdCPU{
		cpu: cpu{
			manufacturer: "Amd",
		},
	}
}

func (a *amd) MakeGPU() iGPU {
	return &amdGPU{
		gpu: gpu{
			manufacturer: "Amd",
		},
	}
}

func (a *amd) MakeMOBO() iMOBO {
	return &amdMOBO{
		mobo: mobo{
			manufacturer: "Amd",
		},
	}
}

func (a *amd) MakeRAM() iRAM {
	return &amdRAM{
		ram: ram{
			manufacturer: "Amd",
		},
	}
}

func (a *amd) MakeDRIVE() iDRIVE {
	return &amdDRIVE{
		drive: drive{
			manufacturer: "Amd",
		},
	}
}

func (a *amd) MakePSU() iPSU {
	return &amdPSU{
		psu: psu{
			manufacturer: "Amd",
		},
	}
}

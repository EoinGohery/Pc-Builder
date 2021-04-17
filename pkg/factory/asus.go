package factory

type asus struct {
}

func (a *asus) MakeCPU() iCPU {
	return &asusCPU{
		cpu: cpu{
			manufacturer: "Asus",
		},
	}
}

func (a *asus) MakeGPU() iGPU {
	return &asusGPU{
		gpu: gpu{
			manufacturer: "Asus",
		},
	}
}

func (a *asus) MakeMOBO() iMOBO {
	return &asusMOBO{
		mobo: mobo{
			manufacturer: "Asus",
		},
	}
}

func (a *asus) MakeRAM() iRAM {
	return &asusRAM{
		ram: ram{
			manufacturer: "Asus",
		},
	}
}

func (a *asus) MakeDRIVE() iDRIVE {
	return &asusDRIVE{
		drive: drive{
			manufacturer: "Asus",
		},
	}
}

func (a *asus) MakePSU() iPSU {
	return &asusPSU{
		psu: psu{
			manufacturer: "Asus",
		},
	}
}

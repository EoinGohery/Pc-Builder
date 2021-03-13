package factory

type intel struct {
}

func (i *intel) makeCPU() iCPU {
	return &intelCPU{
		cpu: cpu{
			manufacturer: "Intel",
		},
	}
}

func (i *intel) makeGPU() iGPU {
	return &intelGPU{
		gpu: gpu{
			manufacturer: "Intel",
		},
	}
}

func (i *intel) makeMOBO() iMOBO {
	return &intelMOBO{
		mobo: mobo{
			manufacturer: "Intel",
		},
	}
}

func (i *intel) makeRAM() iRAM {
	return &intelRAM{
		ram: ram{
			manufacturer: "Intel",
		},
	}
}

func (i *intel) makeDRIVE() iDRIVE {
	return &intelDRIVE{
		drive: drive{
			manufacturer: "Intel",
		},
	}
}

func (i *intel) makePSU() iPSU {
	return &intelPSU{
		psu: psu{
			manufacturer: "Intel",
		},
	}
}

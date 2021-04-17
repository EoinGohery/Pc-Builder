package factory

type intel struct {
}

func (i *intel) MakeCPU() iCPU {
	return &intelCPU{
		cpu: cpu{
			manufacturer: "Intel",
		},
	}
}

func (i *intel) MakeGPU() iGPU {
	return &intelGPU{
		gpu: gpu{
			manufacturer: "Intel",
		},
	}
}

func (i *intel) MakeMOBO() iMOBO {
	return &intelMOBO{
		mobo: mobo{
			manufacturer: "Intel",
		},
	}
}

func (i *intel) MakeRAM() iRAM {
	return &intelRAM{
		ram: ram{
			manufacturer: "Intel",
		},
	}
}

func (i *intel) MakeDRIVE() iDRIVE {
	return &intelDRIVE{
		drive: drive{
			manufacturer: "Intel",
		},
	}
}

func (i *intel) MakePSU() iPSU {
	return &intelPSU{
		psu: psu{
			manufacturer: "Intel",
		},
	}
}

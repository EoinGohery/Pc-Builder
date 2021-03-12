package factory

type intel struct {
}

func (i *intel) makeCPU() iCPU {
	return &intelCPU{
		cpu: cpu{
			name: "Ryzen",
		},
	}
}

func (i *intel) makeGPU() iGPU {
	return &intelGPU{
		gpu: gpu{
			name: "Ryzen",
		},
	}
}

func (i *intel) makeMOBO() iMOBO {
	return &intelMOBO{
		mobo: mobo{
			name: "Ryzen",
		},
	}
}

func (i *intel) makeRAM() iRAM {
	return &intelRAM{
		ram: ram{
			name: "Ryzen",
		},
	}
}

func (i *intel) makeDRIVE() iDRIVE {
	return &intelDRIVE{
		drive: drive{
			name: "Ryzen",
		},
	}
}

func (i *intel) makePSU() iPSU {
	return &intelPSU{
		psu: psu{
			name: "Ryzen",
		},
	}
}

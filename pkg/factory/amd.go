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

func (a *amd) makeMobo() iMOBO {
	return &amdMOBO{
		mobo: mobo{
			name: "Ryzen",
		},
	}
}

func (a *amd) makeRam() iRam {
	return &amdRam{
		ram: ram{
			name: "Ryzen",
		},
	}
}

func (a *amd) makeHardDrive() iDrive {
	return &amdDrive{
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

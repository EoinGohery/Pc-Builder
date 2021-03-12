package factory

type nvidia struct {
}

func (n *nvidia) makeCPU() iCPU {
	return &nvidiaCPU{
		cpu: cpu{
			name: "Ryzen",
		},
	}
}

func (n *nvidia) makeGPU() iGPU {
	return &nvidiaGPU{
		gpu: gpu{
			name: "Ryzen",
		},
	}
}

func (n *nvidia) makeMOBO() iMOBO {
	return &nvidiaMOBO{
		mobo: mobo{
			name: "Ryzen",
		},
	}
}

func (n *nvidia) makeRAM() iRAM {
	return &nvidiaRAM{
		ram: ram{
			name: "Ryzen",
		},
	}
}

func (n *nvidia) makeDRIVE() iDRIVE {
	return &nvidiaDRIVE{
		drive: drive{
			name: "Ryzen",
		},
	}
}

func (n *nvidia) makePSU() iPSU {
	return &nvidiaPSU{
		psu: psu{
			name: "Ryzen",
		},
	}
}

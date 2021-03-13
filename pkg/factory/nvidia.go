package factory

type nvidia struct {
}

func (n *nvidia) makeCPU() iCPU {
	return &nvidiaCPU{
		cpu: cpu{
			manufacturer: "nVidia",
		},
	}
}

func (n *nvidia) makeGPU() iGPU {
	return &nvidiaGPU{
		gpu: gpu{
			manufacturer: "nVidia",
		},
	}
}

func (n *nvidia) makeMOBO() iMOBO {
	return &nvidiaMOBO{
		mobo: mobo{
			manufacturer: "nVidia",
		},
	}
}

func (n *nvidia) makeRAM() iRAM {
	return &nvidiaRAM{
		ram: ram{
			manufacturer: "nVidia",
		},
	}
}

func (n *nvidia) makeDRIVE() iDRIVE {
	return &nvidiaDRIVE{
		drive: drive{
			manufacturer: "nVidia",
		},
	}
}

func (n *nvidia) makePSU() iPSU {
	return &nvidiaPSU{
		psu: psu{
			manufacturer: "nVidia",
		},
	}
}

package factory

type nvidia struct {
}

func (n *nvidia) MakeCPU() iCPU {
	return &nvidiaCPU{
		cpu: cpu{
			manufacturer: "nVidia",
		},
	}
}

func (n *nvidia) MakeGPU() iGPU {
	return &nvidiaGPU{
		gpu: gpu{
			manufacturer: "nVidia",
		},
	}
}

func (n *nvidia) MakeMOBO() iMOBO {
	return &nvidiaMOBO{
		mobo: mobo{
			manufacturer: "nVidia",
		},
	}
}

func (n *nvidia) MakeRAM() iRAM {
	return &nvidiaRAM{
		ram: ram{
			manufacturer: "nVidia",
		},
	}
}

func (n *nvidia) MakeDRIVE() iDRIVE {
	return &nvidiaDRIVE{
		drive: drive{
			manufacturer: "nVidia",
		},
	}
}

func (n *nvidia) MakePSU() iPSU {
	return &nvidiaPSU{
		psu: psu{
			manufacturer: "nVidia",
		},
	}
}

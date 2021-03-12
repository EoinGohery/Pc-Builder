package factory

type seagate struct {
}

func (s *seagate) makeCPU() iCPU {
	return &seagateCPU{
		cpu: cpu{
			name: "Ryzen",
		},
	}
}

func (s *seagate) makeGPU() iGPU {
	return &seagateGPU{
		gpu: gpu{
			name: "Ryzen",
		},
	}
}

func (s *seagate) makeMOBO() iMOBO {
	return &seagateMOBO{
		mobo: mobo{
			name: "Ryzen",
		},
	}
}

func (s *seagate) makeRAM() iRAM {
	return &seagateRAM{
		ram: ram{
			name: "Ryzen",
		},
	}
}

func (s *seagate) makeDRIVE() iDRIVE {
	return &seagateDRIVE{
		drive: drive{
			name: "Ryzen",
		},
	}
}

func (s *seagate) makePSU() iPSU {
	return &seagatePSU{
		psu: psu{
			name: "Ryzen",
		},
	}
}

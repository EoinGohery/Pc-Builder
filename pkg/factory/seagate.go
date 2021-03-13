package factory

type seagate struct {
}

func (s *seagate) makeCPU() iCPU {
	return &seagateCPU{
		cpu: cpu{
			manufacturer: "Seagate",
		},
	}
}

func (s *seagate) makeGPU() iGPU {
	return &seagateGPU{
		gpu: gpu{
			manufacturer: "Seagate",
		},
	}
}

func (s *seagate) makeMOBO() iMOBO {
	return &seagateMOBO{
		mobo: mobo{
			manufacturer: "Seagate",
		},
	}
}

func (s *seagate) makeRAM() iRAM {
	return &seagateRAM{
		ram: ram{
			manufacturer: "Seagate",
		},
	}
}

func (s *seagate) makeDRIVE() iDRIVE {
	return &seagateDRIVE{
		drive: drive{
			manufacturer: "Seagate",
		},
	}
}

func (s *seagate) makePSU() iPSU {
	return &seagatePSU{
		psu: psu{
			manufacturer: "Seagate",
		},
	}
}

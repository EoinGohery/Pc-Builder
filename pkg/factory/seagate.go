package factory

type seagate struct {
}

func (s *seagate) MakeCPU() iCPU {
	return &seagateCPU{
		cpu: cpu{
			manufacturer: "Seagate",
		},
	}
}

func (s *seagate) MakeGPU() iGPU {
	return &seagateGPU{
		gpu: gpu{
			manufacturer: "Seagate",
		},
	}
}

func (s *seagate) MakeMOBO() iMOBO {
	return &seagateMOBO{
		mobo: mobo{
			manufacturer: "Seagate",
		},
	}
}

func (s *seagate) MakeRAM() iRAM {
	return &seagateRAM{
		ram: ram{
			manufacturer: "Seagate",
		},
	}
}

func (s *seagate) MakeDRIVE() iDRIVE {
	return &seagateDRIVE{
		drive: drive{
			manufacturer: "Seagate",
		},
	}
}

func (s *seagate) MakePSU() iPSU {
	return &seagatePSU{
		psu: psu{
			manufacturer: "Seagate",
		},
	}
}

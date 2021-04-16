package builder

type cpuBuilder struct {
	url []byte
}

func newCpuBuilder() *cpuBuilder {
	return &cpuBuilder{}
}

func (b *cpuBuilder) sendRequest(v visitor) {
	v.visitForCPU(b)
}

func (b *cpuBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

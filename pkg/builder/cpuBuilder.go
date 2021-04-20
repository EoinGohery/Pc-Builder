package builder

type cpuBuilder struct {
	url []byte
}

//connstructor
func newCpuBuilder() *cpuBuilder {
	return &cpuBuilder{}
}

//send request that triggers the vistor method fort this object type
func (b *cpuBuilder) sendRequest(v visitor) {
	v.visitForCPU(b)
}

//get teh result of teh send request.
func (b *cpuBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

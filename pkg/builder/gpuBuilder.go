package builder

type gpuBuilder struct {
	url []byte
}

//connstructor
func newGpuBuilder() *gpuBuilder {
	return &gpuBuilder{}
}

//send request that triggers the vistor method fort this object type
func (b *gpuBuilder) sendRequest(v visitor) {
	v.visitForGPU(b)
}

//get teh result of teh send request.
func (b *gpuBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

package builder

type gpuBuilder struct {
	url []byte
}

func newGpuBuilder() *gpuBuilder {
	return &gpuBuilder{}
}

func (b *gpuBuilder) sendRequest(v visitor) {
	v.visitForGPU(b)
}

func (b *gpuBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

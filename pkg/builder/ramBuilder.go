package builder

type ramBuilder struct {
	url []byte
}

func newRamBuilder() *ramBuilder {
	return &ramBuilder{}
}

func (b *ramBuilder) sendRequest(v visitor) {
	v.visitForRam(b)
}

func (b *ramBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

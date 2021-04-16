package builder

type psuBuilder struct {
	url []byte
}

func newPsBuilder() *psuBuilder {
	return &psuBuilder{}
}

func (b *psuBuilder) sendRequest(v visitor) {
	v.visitForPSU(b)
}

func (b *psuBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

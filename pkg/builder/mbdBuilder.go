package builder

type mbdBuilder struct {
	url []byte
}

func newMbdBuilder() *mbdBuilder {
	return &mbdBuilder{}
}

func (b *mbdBuilder) sendRequest(v visitor) {
	v.visitForMB(b)
}

func (b *mbdBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

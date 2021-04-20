package builder

type psuBuilder struct {
	url []byte
}

//connstructor
func newPsBuilder() *psuBuilder {
	return &psuBuilder{}
}

//send request that triggers the vistor method fort this object type
func (b *psuBuilder) sendRequest(v visitor) {
	v.visitForPSU(b)
}

//get teh result of teh send request.
func (b *psuBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

package builder

type ramBuilder struct {
	url []byte
}

//constuctor
func newRamBuilder() *ramBuilder {
	return &ramBuilder{}
}

//send request that triggers the vistor method fort this object type
func (b *ramBuilder) sendRequest(v visitor) {
	v.visitForRam(b)
}

//get teh result of teh send request.
func (b *ramBuilder) getRequest() request {
	return request{
		url: b.url,
	}
}

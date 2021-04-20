package builder

type mbdBuilder struct {
	url []byte
}

//connstructor
func newMbdBuilder() *mbdBuilder {
	return &mbdBuilder{}
}

//send request that triggers the vistor method fort this object type
func (b *mbdBuilder) sendRequest(v visitor) {
	v.visitForMB(b)
}

//get teh result of teh send request.
func (b *mbdBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

package builder

type driveBuilder struct {
	url []byte
}

//connstructor
func newDriveBuilder() *driveBuilder {
	return &driveBuilder{}
}

//send request that triggers the vistor method fort this object type
func (b *driveBuilder) sendRequest(v visitor) {
	v.visitForDrive(b)
}

//get teh result of teh send request.
func (b *driveBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

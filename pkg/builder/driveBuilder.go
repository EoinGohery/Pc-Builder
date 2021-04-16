package builder

type driveBuilder struct {
	url []byte
}

func newDriveBuilder() *driveBuilder {
	return &driveBuilder{}
}

func (b *driveBuilder) sendRequest(v visitor) {
	v.visitForDrive(b)
}

func (b *driveBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

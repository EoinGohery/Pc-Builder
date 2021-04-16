package builder

type director struct {
	builder iBuilder
}

func NewDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) BuildRequest() request {
	d.builder.sendRequest(&send{})
	return d.builder.getRequest()
}

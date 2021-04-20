package builder

type director struct {
	builder iBuilder
}

//director that builds the required objecvt from their respective builder.
func NewDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

//sets the builder type.
func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

//triggers the send reuest, which triggers the visotr. Then get teh returned data.
func (d *director) BuildRequest() request {
	d.builder.sendRequest(&send{})
	return d.builder.getRequest()
}

package builder

type iBuilder interface {
	sendRequest(v visitor)
	getRequest() request
}

func GetBuilder(builderType string) iBuilder {
	if builderType == "cpu" {
		return &cpuBuilder{}
	}

	if builderType == "gpu" {
		return &gpuBuilder{}
	}

	if builderType == "driver" {
		return &driveBuilder{}
	}

	if builderType == "mbd" {
		return &mbdBuilder{}
	}

	if builderType == "ps" {
		return &psuBuilder{}
	}

	if builderType == "ram" {
		return &ramBuilder{}
	}
	return nil
}

package builder

type iBuilder interface {
	sendRequest()
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
		return &driverBuilder{}
	}

	if builderType == "mbd" {
		return &mbdBuilder{}
	}

	if builderType == "ps" {
		return &psBuilder{}
	}

	if builderType == "ram" {
		return &ramBuilder{}
	}
	return nil
}

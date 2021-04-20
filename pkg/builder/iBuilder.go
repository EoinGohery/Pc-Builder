package builder

type iBuilder interface {
	sendRequest(v visitor)
	getRequest() request
}

//create the builder othe correct type
func GetBuilder(builderType string) iBuilder {
	if builderType == "cpu" {
		return &cpuBuilder{}
	}

	if builderType == "gpu" {
		return &gpuBuilder{}
	}

	if builderType == "drive" {
		return &driveBuilder{}
	}

	if builderType == "mbd" {
		return &mbdBuilder{}
	}

	if builderType == "psu" {
		return &psuBuilder{}
	}

	if builderType == "ram" {
		return &ramBuilder{}
	}
	return nil
}

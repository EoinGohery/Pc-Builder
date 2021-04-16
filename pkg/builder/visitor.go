package builder

type visitor interface {
	visitForCPU(*cpuBuilder)
	visitForGPU(*gpuBuilder)
	visitForDrive(*driveBuilder)
	visitForPSU(*psuBuilder)
	visitForRam(*ramBuilder)
	visitForMB(*mbdBuilder)
}

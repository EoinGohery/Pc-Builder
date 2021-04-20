package builder

//the visitor in allows for objects to call teh methods in the visitor send.go class
type visitor interface {
	visitForCPU(*cpuBuilder)
	visitForGPU(*gpuBuilder)
	visitForDrive(*driveBuilder)
	visitForPSU(*psuBuilder)
	visitForRam(*ramBuilder)
	visitForMB(*mbdBuilder)
}

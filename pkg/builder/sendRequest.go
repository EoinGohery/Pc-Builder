package builder

import (
	"fmt"
)

//sends the request, triggered by executer interceptor
func FireRequest(input string) string {
	partBuilder := GetBuilder(input)
	director := NewDirector(partBuilder)
	partRequest := director.BuildRequest()
	return fmt.Sprintf("%s", partRequest)
}

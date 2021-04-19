package builder

import (
	"fmt"
)

func FireRequest(input string) string {
	partBuilder := GetBuilder(input)
	director := NewDirector(partBuilder)
	partRequest := director.BuildRequest()
	return fmt.Sprintf("%s", partRequest)
}

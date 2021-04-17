package filter

import (
	"CS4227/pkg/builder"
	"fmt"
)

func SendRequest(input string) string {
	partBuilder := builder.GetBuilder(input)
	director := builder.NewDirector(partBuilder)
	partRequest := director.BuildRequest()
	return fmt.Sprintf("%s", partRequest)
}

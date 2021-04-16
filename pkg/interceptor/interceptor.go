package interceptor

import "fmt"

type Interceptor struct {
}

func (i Interceptor) execute(request string) {
	fmt.Sprintf("in Interceptor: %s", request)
}
package interceptor

import (
	"CS4227/pkg/filter"
)

type Interceptor struct {
}

func (i Interceptor) execute(request string) string {
	return filter.SendRequest(request)
}

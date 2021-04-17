package interceptor

import (
	"CS4227/pkg/filter"
)

type ExecuteInterceptor struct {
}

func (i ExecuteInterceptor) execute(request string) string {
	return filter.SendRequest(request)
}

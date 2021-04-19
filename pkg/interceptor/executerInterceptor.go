package interceptor

import (
	"CS4227/pkg/builder"
)

type ExecuteInterceptor struct {
}

func (i ExecuteInterceptor) execute(request string) string {
	return builder.FireRequest(request)
}

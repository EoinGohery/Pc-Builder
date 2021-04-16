package interceptor

import "fmt"

type LoggingInterceptor struct {

}

func (l LoggingInterceptor) execute(request string) {
	fmt.Sprintf("in LoggingInterceptor: %s", request)
}

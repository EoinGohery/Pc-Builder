package interceptor

import "fmt"

type FilterInterceptor struct {

}

func (f FilterInterceptor) execute(request string) {
	fmt.Sprintf("in FilterInterceptor: %s", request)
}
package interceptor

import "CS4227/pkg/factory"

type InterceptorManager struct {
	executerInterceptor ExecuteInterceptor
	loggingInterceptor  LoggingInterceptor
	filterInterceptor   FilterInterceptor
}

func (i InterceptorManager) SetExecuter(executerInterceptor ExecuteInterceptor) {
	i.executerInterceptor = executerInterceptor
}

func (i InterceptorManager) SetFilter(filterInterceptor FilterInterceptor) {
	i.filterInterceptor = filterInterceptor
}

func (i InterceptorManager) SetLogging(loggingInterceptor LoggingInterceptor) {
	i.loggingInterceptor = loggingInterceptor
}

func (i InterceptorManager) Execute(request string, filter string) []factory.Component {
	i.loggingInterceptor.execute(request, filter)
	returned := i.executerInterceptor.execute(request)
	result := i.filterInterceptor.execute(request, returned, filter)
	return result
}

package interceptor

import "CS4227/pkg/factory"

type InterceptorManager struct {
	command            Interceptor //replace with command object later
	loggingInterceptor LoggingInterceptor
	filterInterceptor  FilterInterceptor
}

func (i InterceptorManager) SetFilter(filterInterceptor FilterInterceptor) {
	i.filterInterceptor = filterInterceptor
}

func (i InterceptorManager) SetLogging(loggingInterceptor LoggingInterceptor) {
	i.loggingInterceptor = loggingInterceptor
}

func (i InterceptorManager) Execute(request string, filter string) []factory.Component {
	//i.loggingInterceptor.execute(request)
	returned := i.command.execute(request)
	result := i.filterInterceptor.execute(request, returned, filter)
	return result
}

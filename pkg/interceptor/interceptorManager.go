package interceptor

type InterceptorManager struct {
	command            Interceptor //replace with command object later
	loggingInterceptor LoggingInterceptor
	filterInterceptor  FilterInterceptor
}

func (i InterceptorManager) setFilter(filterInterceptor FilterInterceptor) {
	i.filterInterceptor = filterInterceptor
}

func (i InterceptorManager) setLogging(loggingInterceptor LoggingInterceptor) {
	i.loggingInterceptor = loggingInterceptor
}

func (i InterceptorManager) execute(request string) {
	i.loggingInterceptor.execute(request)
	i.filterInterceptor.execute(request)
	i.command.execute(request)
}

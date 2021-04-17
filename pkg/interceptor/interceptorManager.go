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

func (i InterceptorManager) execute(request string, filter map[string]interface{}) map[string]interface{} {
	i.loggingInterceptor.execute(request)
	returned := i.command.execute(request)
	result := i.filterInterceptor.execute(returned, filter)
	return result
}

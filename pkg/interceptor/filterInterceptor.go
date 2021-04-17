package interceptor

import (
	"CS4227/pkg/factory"
	"CS4227/pkg/filter"
)

type FilterInterceptor struct {
}

func (f FilterInterceptor) execute(request string, data string, filtertoSet string) []factory.Component {
	formatted := filter.JsonParse(request, data)
	filtered := filter.Filter(formatted, filtertoSet)
	return filtered
}

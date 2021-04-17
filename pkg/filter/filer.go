package filter

import (
	"CS4227/pkg/factory"
)

func Filter(data []factory.Component, filter string) []factory.Component {
	if filter != "" {
		var returned []factory.Component
		for i := range data {
			if data[i].GetFilter() == filter {
				returned = append(returned, data[i])
			}
		}
		return returned
	} else {
		return data
	}
}

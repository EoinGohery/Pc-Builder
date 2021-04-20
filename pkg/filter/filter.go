package filter

import (
	"CS4227/pkg/factory"
)

func Filter(data []factory.Component, filter string) []factory.Component {
	if filter != "" {
		var returned []factory.Component
		var x int = 0
		for i := range data {
			if data[i].GetFilter() == filter {
				x = x + 1
				data[i].SetID(x)
				returned = append(returned, data[i])
			}
		}
		return returned
	} else {
		return data
	}
}

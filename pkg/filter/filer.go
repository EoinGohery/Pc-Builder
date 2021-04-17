package filter

import (
	"CS4227/pkg/factory"
	"fmt"
)

func Filter(data []factory.Component, filter string) []factory.Component {
	var returned []factory.Component
	for i := range data {
		fmt.Println(data[i])
		if data[i].GetFilter() == filter {
			returned = append(returned, data[i])
		}
	}
	return returned
}

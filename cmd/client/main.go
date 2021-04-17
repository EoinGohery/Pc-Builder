package client

import (
	"CS4227/pkg/factory"
	"CS4227/pkg/interceptor"
	"fmt"
)

//Run client
func Run() {

	//factory.PrintDetails()

	// fmt.Println("\nEnter one of the following (1) cpu (2) gpu (3) driver (4) mbd (5) ps (6) ram: ")
	// var input string
	// fmt.Scanln(&input)

	var manager interceptor.InterceptorManager
	var logger interceptor.LoggingInterceptor
	var filterer interceptor.FilterInterceptor

	manager.SetFilter(filterer)
	manager.SetLogging(logger)

	all := manager.Execute("cpu", "AM4")

	for i := range all {
		fmt.Print(all[i].ToString())
	}

	tower := &factory.Tower{}
	tower.Add(all[0])
	tower.Print()
}

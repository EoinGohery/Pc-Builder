package client

import (
	"CS4227/pkg/filter"
	"fmt"
)

//Run client
func Run() {

	//factory.PrintDetails()

	// fmt.Println("\nEnter one of the following (1) cpu (2) gpu (3) driver (4) mbd (5) ps (6) ram: ")
	// var input string
	// fmt.Scanln(&input)

	all := filter.SendRequest("gpu")

	fmt.Print(all)
}

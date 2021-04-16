package client

import (
	"CS4227/pkg/builder"
	"fmt"
)

//Run client
func Run() {

	//factory.PrintDetails()

	fmt.Println("\nEnter one of the following (1) cpu (2) gpu (3) driver (4) mbd (5) ps (6) ram: ")
	var input string
	fmt.Scanln(&input)
	partBuilder := builder.GetBuilder(input)
	director := builder.NewDirector(partBuilder)
	partRequest := director.BuildRequest()

	fmt.Printf("Result: %s\n", partRequest)
}

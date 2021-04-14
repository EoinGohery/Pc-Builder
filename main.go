package main

import (
	"CS4227/cmd/backend"
	"CS4227/cmd/client"
	"fmt"
	"time"
)

func main() {

	for {
		if !backend.Run() {
			break
		}
		fmt.Printf("\nNo server connection detected. Ensure the server has been started.")
		time.Sleep(time.Second)
	}
	client.Run()
}

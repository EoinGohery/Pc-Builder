package main

import (
	"CS4227/cmd/backend"
	"CS4227/cmd/client"
)

func main() {
	backend.Run()
	client.Run()
}

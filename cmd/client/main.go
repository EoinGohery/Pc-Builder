package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Run client
func Run() {

	//factory.PrintDetails()

	fmt.Println("\nEnter number (1) Cpus (2) Gpus (3) Drivers (4) Motherboards (5) PSU (6) Ram: ")
	var input int
	fmt.Scanln(&input)
	switch input {
	case 1:
		resp, err := http.Get("http://localhost:8080/api/parts/getCPUs")
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s\n", body)
	case 2:
		resp, err := http.Get("http://localhost:8080/api/parts/getGPUs")
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s\n", body)
	case 3:
		resp, err := http.Get("http://localhost:8080/api/parts/getDrivers")
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s\n", body)
	case 4:
		resp, err := http.Get("http://localhost:8080/api/parts/getMBDs")
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s\n", body)
	case 5:
		resp, err := http.Get("http://localhost:8080/api/parts/getPS")
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s\n", body)
	case 6:
		resp, err := http.Get("http://localhost:8080/api/parts/getRam")
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s\n", body)
	}
}

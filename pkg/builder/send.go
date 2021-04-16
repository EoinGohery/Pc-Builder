package builder

import (
	"io/ioutil"
	"log"
	"net/http"
)

type send struct {
}

func (s *send) visitForCPU(b *cpuBuilder) {
	resp, err := http.Get("http://localhost:8080/api/parts/getCPUs")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Printf("%s\n", body)
	b.url = body
}

func (s *send) visitForDrive(b *driveBuilder) {
	resp, err := http.Get("http://localhost:8080/api/parts/getDrivers")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Printf("%s\n", body)
	b.url = body
}

func (s *send) visitForGPU(b *gpuBuilder) {
	resp, err := http.Get("http://localhost:8080/api/parts/getGPUs")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Printf("%s\n", body)
	b.url = body
}

func (s *send) visitForPSU(b *psuBuilder) {
	resp, err := http.Get("http://localhost:8080/api/parts/getPS")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Printf("%s\n", body)
	b.url = body
}

func (s *send) visitForMB(b *mbdBuilder) {
	resp, err := http.Get("http://localhost:8080/api/parts/getMBDs")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Printf("%s\n", body)
	b.url = body
}

func (s *send) visitForRam(b *ramBuilder) {
	resp, err := http.Get("http://localhost:8080/api/parts/getRam")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Printf("%s\n", body)
	b.url = body
}

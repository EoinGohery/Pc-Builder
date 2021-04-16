package builder

import (
	"io/ioutil"
	"log"
	"net/http"
)

type driverBuilder struct {
	url []byte
}

func newDriverBuilder() *driverBuilder {
	return &driverBuilder{}
}

func (b *driverBuilder) sendRequest() {
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

func (b *driverBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

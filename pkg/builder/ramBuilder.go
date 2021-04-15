package builder

import (
	"io/ioutil"
	"log"
	"net/http"
)

type ramBuilder struct {
	url []byte
}

func newRamBuilder() *ramBuilder {
	return &ramBuilder{}
}

func (b *ramBuilder) sendRequest() {
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

func (b *ramBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}

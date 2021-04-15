package builder

import (
	"io/ioutil"
	"log"
	"net/http"
)

type mbdBuilder struct {
	url []byte
}

func newMbdBuilder() *mbdBuilder {
	return &mbdBuilder{}
}

func (b *mbdBuilder) sendRequest() {
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

func (b *mbdBuilder) getRequest() request {
	return request{
		// extension: b.extension,
		url: b.url,
	}
}
